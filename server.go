package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Hw5_GoAuctionSystem/proto"
	"google.golang.org/grpc"
)

// Representation of a server.

var (
	serverId int32 // Server's ID.
	bidCount int   // Count of the total amount of bids.
	bidRnd   int   // bidding round.
	bidVal   int   // Value of current bid.
	bidOver  bool  // If a bid is over or not.
)

// stopRound | stops the round.
func stopRound(id int32) {
	log.Printf("Bidding round %d over: winner %v, amount %d, total bids %d", bidRnd+1, id, bidVal, bidCount)

	bidOver = true

	time.Sleep(time.Second * 5)

	initiateRound()
}

func initiateRound() {
	bidCount = 0
	bidVal = 0
	bidRnd++
	bidOver = false
	log.Printf("Starting bidding round %d, starting amount %d", bidRnd, bidVal)
}

// Bid | handles bidding.
func (s *server) Bid(_ context.Context, bid *GoAuctionSystem.BidPost) (*GoAuctionSystem.Ack, error) {
	log.Printf("Client %v bid amount %d", bid.Id, bid.Amount)

	if int(bid.Amount) > bidVal {
		bidVal = int(bid.Amount)

		if bidCount++; bidCount == 10 {
			go stopRound(bid.Id)
		}

		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_SUCCESS}, nil

	} else if int(bid.Amount) < bidVal {

		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_FAIL}, nil

	} else {

		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_EXCEPTION}, nil

	}
}

type server struct {
	GoAuctionSystem.UnimplementedAuctionSystemServer
}

// Result | returns the server's highest bid received and informs whether round is over or not.
func (s *server) Result(context.Context, *GoAuctionSystem.Empty) (*GoAuctionSystem.Outcome, error) {
	// ??
	return &GoAuctionSystem.Outcome{Amount: int32(bidVal), Over: bidOver}, nil
}

// StartServer | Sets up and starts the server.
func StartServer() {
	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", serverId+5000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register a new gRPC server as an auction system server.
	s := grpc.NewServer()
	GoAuctionSystem.RegisterAuctionSystemServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	// Serve at address
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	// "Shakes the bag" to ensure randomization.
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:] // args: <port number>
	portId, _ := strconv.ParseInt(args[0], 10, 32)
	serverId = int32(portId)

	log.Printf("Starting bidding round %d, starting amount %d", 1, bidVal)
	StartServer()
}
