package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Hw5_GoAuctionSystem/proto"
	"google.golang.org/grpc"
)

var (
	cid         int32
	lamport     int64
	serverCount int
	currentBid  int
	roundOver   bool
	chanDone    []chan bool
	servers     []GoAuctionSystem.AuctionSystemClient
)

// isServerAlive | checks whether a server responds or not.
func isServerAlive(err error, serverId int) bool {
	if err != nil {
		log.Printf("Server %v unresponsive, connection disconnected...", serverId)
		servers[serverId] = nil
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

	for i := 0; i < serverCount; i++ {
		if servers[i] != nil {
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
