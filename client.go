// --------------------------- //
// ---------- IMPORT --------- //
// --------------------------- //
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
    "math/rand"

	"github.com/Hw5_GoAuctionSystem/proto"
	"google.golang.org/grpc"
)

// --------------------------- //
// --------- GLOBALS --------- //
// --------------------------- //
var (
	id          int32
	lamport     int64
	serverCount int
	currentBid  int
    roundOver   bool
    chanDone    []chan bool
	servers     []GoAuctionSystem.AuctionSystemClient
)

// --------------------------- //
// ---------- CLIENT --------- //
// --------------------------- //
func CheckServer(err error, serverId int) bool {
    if err != nil {
        log.Printf("Server %v unresponsive, connection disconnected...", serverId)
        servers[serverId] = nil
        chanDone[serverId] <- true
        return false
    }
    return true
}

func BroadcastBid(amount int32) {
    lamport++
    timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
    currBid := GoAuctionSystem.BidPost{Id: id, Amount: amount, Lamport: lamport}
    log.Printf("Bidding amount %d", amount)
    for i, s := range servers {
        if s != nil {
            ack, err := s.Bid(timeout, &currBid)
            if CheckServer(err,i) {
                switch ack.Ack {
                    case GoAuctionSystem.Acks_ACK_FAIL:
                        log.Printf("Bidding server %v failed!",i)
                    case GoAuctionSystem.Acks_ACK_SUCCESS:
                        log.Printf("Bidding server %v sucess!",i)
                    case GoAuctionSystem.Acks_ACK_EXCEPTION:
                        log.Printf("Bidding server %v exception",i)
                }
            }
        }
    }
}

func GetResult() *GoAuctionSystem.Outcome {
    lamport++
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
    for i := 0; i < serverCount; i++ {
        if servers[i] != nil {
	        result, err := servers[i].Result(timeout, &GoAuctionSystem.Empty{})
	        if !CheckServer(err,i) {
                continue
            }
            return result
        }
    }
    return nil
}

func FrontEnd(servers int) {
	for {
        if serverCount == servers {
            result := GetResult()

		    if result.Over {
                if !roundOver {
                    roundOver = true
                    currentBid = 0
                    log.Printf("Round over, total bidding amount: %v", result.Amount)
                    return
                }
                time.Sleep(time.Second*time.Duration(rand.Intn(2)+1))
                continue
		    }

            roundOver = false
            BroadcastBid(result.Amount+int32(rand.Intn(500)))
            time.Sleep(time.Second*time.Duration(rand.Intn(2)+1))
        }
	}
}

func DialServer(serverId int) {
	// Dial server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serverId+5000), grpc.WithInsecure())
    if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Setup client
	servers[serverId] = GoAuctionSystem.NewAuctionSystemClient(conn)
    chanDone[serverId] = make(chan bool)
	log.Printf("Client connected to server...")
    serverCount++

	// Closes connection
	<-chanDone[serverId]
    conn.Close()
}

// --------------------------- //
// ---------- SETUP ---------- //
// --------------------------- //
func main() {
    rand.Seed(time.Now().UnixNano())
	args := os.Args[1:] // args: <client ID> <server Count>
	aid, _ := strconv.ParseInt(args[0], 10, 32)
    sc, _ := strconv.ParseInt(args[1], 10, 32)
    br, _ := strconv.ParseInt(args[2], 10, 32)
	id = int32(aid)
    chanDone = make([]chan bool, int(sc))
    servers = make([]GoAuctionSystem.AuctionSystemClient, int(sc))

	for i := 0; i < int(sc); i++ {
		go DialServer(i)
	}

	for i := 0; i < int(br); i++ {
        log.Printf("Starting bidding round %d/%d", i+1, int(br))
		FrontEnd(int(sc))
	}
}
