package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Hw5_GoAuctionSystem/proto"
	"google.golang.org/grpc"
)

var (
	cid        int32
	lamport    int64
	currentBid int
	roundOver  bool
	chanDone   []chan bool
	servers    map[int]GoAuctionSystem.AuctionSystemClient
)

// isServerAlive | checks whether a server responds or not.
func isServerAlive(err error, serverId int) bool {
	if err != nil {
		log.Printf("Server %v unresponsive, connection disconnected...", serverId)
		delete(servers, serverId)
		chanDone[serverId] <- true
		return false
	}
	return true
}

// BroadcastBid | Announces and posts the bid to the server(s).
func BroadcastBid(amount int32) {
	lamport++

	// timeout checks whether server responds within timeout. Otherwise, assume it's dead.
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)

	currBid := GoAuctionSystem.BidPost{Id: cid, Amount: amount, Lamport: lamport}

	log.Printf("Bidding amount: %d", amount)

	// checks if the server(s) are/is alive.
	// Handles then acknowledgements.
	for i, s := range servers {
		if s != nil {
			ack, err := s.Bid(timeout, &currBid)
			if isServerAlive(err, i) {
				switch ack.Ack {
				case GoAuctionSystem.Acks_ACK_FAIL:
					log.Printf("Bidding server %v failed!", i)
				case GoAuctionSystem.Acks_ACK_SUCCESS:
					log.Printf("Bidding server %v sucess!", i)
				case GoAuctionSystem.Acks_ACK_EXCEPTION:
					log.Printf("Bidding server %v exception", i)
				}
			}
		}
	}
}

// gets the final result of the bidding.
func GetResult() *GoAuctionSystem.Outcome {

	lamport++
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)

	for i := 0; i < len(servers); i++ {
		//if servers[i] != nil {
		if _, serverExists := servers[i]; serverExists {
			result, err := servers[i].Result(timeout, &GoAuctionSystem.Empty{})

			if !isServerAlive(err, i) {
				continue
			}

			return result
		}
	}

	return nil
}

// FrontEnd || Middleman interconnect client to servers.
func FrontEnd(reqServerAmount int) {
	for {
		if len(servers) == reqServerAmount {
			result := GetResult()

			if result.Over {
				if !roundOver {
					roundOver = true
					currentBid = 0
					log.Printf("Round over, total bidding amount: %v", result.Amount)
					return
				}
				time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
				continue
			}

			roundOver = false
			BroadcastBid(result.Amount + int32(rand.Intn(500)))
			time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
		}
	}
}

func DialServer(serverId int) {
	// Dial server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serverId+5000), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Setup client
	servers[serverId] = GoAuctionSystem.NewAuctionSystemClient(conn)
	chanDone[serverId] = make(chan bool)
	log.Printf("Client connected to server...")

	// Closes connection
	<-chanDone[serverId]
	conn.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:] // args: <client ID> <server Count>
	aid, _ := strconv.ParseInt(args[0], 10, 32)
	sc, _ := strconv.ParseInt(args[1], 10, 32)
	br, _ := strconv.ParseInt(args[2], 10, 32)
	cid = int32(aid)

	servers = make(map[int]GoAuctionSystem.AuctionSystemClient)
	chanDone = make([]chan bool, int(sc))

	for i := 0; i < int(sc); i++ {
		go DialServer(i)
	}

	for i := 0; i < int(br); i++ {
		log.Printf("Starting bidding round %d/%d", i+1, int(br))
		FrontEnd(int(sc))
	}
}
