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
	sid      int32
	bidCount int
	bidRnd   int
	bidVal   int
	bidOver  bool
)

type server struct {
	GoAuctionSystem.UnimplementedAuctionSystemServer
}

func BidBreak(id int32) {
	if bidCount++; bidCount == 10 {
		log.Printf("Bidding round %d over: winner %v, amount %d, total bids %d", bidRnd+1, id, bidVal, bidCount)
		bidOver = true
		time.Sleep(time.Second * 5)
		bidCount = 0
		bidVal = 0
		bidRnd++
		bidOver = false
		log.Printf("Starting bidding round %d, starting amount %d", bidRnd, bidVal)
	}
}

func (s *server) Bid(context context.Context, bid *GoAuctionSystem.BidPost) (*GoAuctionSystem.Ack, error) {
	log.Printf("Client %v bid amount %d", bid.Id, bid.Amount)

	if int(bid.Amount) > bidVal {
		bidVal = int(bid.Amount)
		go BidBreak(bid.Id)
		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_SUCCESS}, nil
	} else if int(bid.Amount) < bidVal {
		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_FAIL}, nil
	} else {
		return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_EXCEPTION}, nil
	}
}

func (s *server) Result(context.Context, *GoAuctionSystem.Empty) (*GoAuctionSystem.Outcome, error) {
	// ??
	return &GoAuctionSystem.Outcome{Amount: int32(bidVal), Over: bidOver}, nil
}

func StartServer() {
	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", sid+5000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register grpc server
	s := grpc.NewServer()
	GoAuctionSystem.RegisterAuctionSystemServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())

	// Serve at addresse
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// --------------------------- //
// ---------- SETUP ---------- //
// --------------------------- //
func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:] // args: <port number>
	pid, _ := strconv.ParseInt(args[0], 10, 32)
	sid = int32(pid)

	log.Printf("Starting bidding round %d, starting amount %d", 1, bidVal)
	StartServer()
}
