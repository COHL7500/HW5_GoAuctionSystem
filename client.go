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

// Representation of a client.

var (
	clientId     int32 // Client's ID.
	lamport      int64 // Lamport to ensure sequentialism - Happens-before causality.
	roundOver    bool
	isServerDone []chan bool
	servers      map[int]GoAuctionSystem.AuctionSystemClient
)

// isServerAlive | checks whether a server responds or not.
func isServerAlive(err error, serverId int) bool {
	if err != nil {
		log.Printf("Server %v unresponsive, connection disconnected...", serverId)
		delete(servers, serverId)
		isServerDone[serverId] <- true
		return false
	}
	return true
}

// BroadcastBid | Announces and posts the bid to the server(s).
func BroadcastBid(amount int32) {
	lamport++

	// timeout checks whether server responds within timeout. Otherwise, assume it's dead.
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)

	currBid := GoAuctionSystem.BidPost{Id: clientId, Amount: amount, Lamport: lamport}

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

// GetResult | returns the current highest bid of the bidding by calling all the servers' Result function.
func GetResult() *GoAuctionSystem.Outcome {

	lamport++
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)

	//
	for i := 0; i < len(servers); i++ {

		// if servers[i] returns a server, serverExists is true. Otherwise, return nil.
		if _, serverExists := servers[i]; serverExists {

			// Initiates result variable with the final bid of the bidding.
			result, err := servers[i].Result(timeout, &GoAuctionSystem.Empty{})

			// Skips the current server if it has become unresponsive after the request.
			// NOTE: isServerAlive and serverExists is different:
			if !isServerAlive(err, i) {
				continue
			}

			return result
		}
	}

	return nil
}

// FrontEnd || Middleman interconnect client to servers.
// We simulate a client interacting with the bidding.
func FrontEnd(reqServerAmount int) {
	// While bidding isn't over...
	for {

		// Checks if the amount of servers corresponds to the requested amount of servers.
		if len(servers) == reqServerAmount {
			result := GetResult()

			// If bid is over, check results. Otherwise, bid over the current bid.
			if result.Over {
				if !roundOver {
					roundOver = true
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

// DialServer | initialize and connect to a server.
func DialServer(serverId int) {
	// Dial server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serverId+5000), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Setup client
	servers[serverId] = GoAuctionSystem.NewAuctionSystemClient(conn)
	isServerDone[serverId] = make(chan bool)
	log.Printf("Client connected to server %v", serverId)

	// If a value is not received from isServerDone, close connection. Otherwise, close connection.
	<-isServerDone[serverId]
	errClose := conn.Close()
	if errClose != nil {
		log.Fatalf("Error occured upon closing connection: %v", errClose)
	}
}

func main() {

	// Randomizes the seed of our randomizer. This is equivalent to shake a bag to randomize the outcome.
	rand.Seed(time.Now().UnixNano())

	// arguments given from the clinet in console for setting up client.
	args := os.Args[1:] // args: <client ID> <server Count> < bidding Rounds >
	id, _ := strconv.ParseInt(args[0], 10, 32)
	serverAmount, _ := strconv.ParseInt(args[1], 10, 32)
	bidRound, _ := strconv.ParseInt(args[2], 10, 32)
	clientId = int32(id)

	// Initiate a hashmap consisting of all servers.
	servers = make(map[int]GoAuctionSystem.AuctionSystemClient)

	// Initiate an array of the channels for every server.
	// We use them to communicate whether a channel is active or not.
	isServerDone = make([]chan bool, int(serverAmount))

	// Dial up each server.
	for i := 0; i < int(serverAmount); i++ {
		go DialServer(i)
	}

	// Run N amount of bidding rounds.
	for i := 0; i < int(bidRound); i++ {
		log.Printf("Starting bidding round %d/%d", i+1, int(bidRound))
		FrontEnd(int(serverAmount))
	}
}
