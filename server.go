// ---------------------------- //
// ---------- IMPORT ---------- //
// ---------------------------- //
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "strconv"

    "github.com/Hw5_GoAuctionSystem/proto"
    "google.golang.org/grpc"
)

// --------------------------- //
// --------- GLOBALS --------- //
// --------------------------- //
var (
	id      int32
	bidRnd  int = 0
	bidVal  int = 0
)

type server struct {
	GoAuctionSystem.UnimplementedAuctionSystemServer
}

// --------------------------- //
// ---------- SERVER --------- //
// --------------------------- //
func (s *server) Bid(context context.Context, bid *GoAuctionSystem.BidPost) (*GoAuctionSystem.Ack, error) {
	log.Printf("Got client %v bid, amount: %v", bid.Id, bid.Amount)
    bidVal = int(bid.Amount)
    return &GoAuctionSystem.Ack{Ack: GoAuctionSystem.Acks_ACK_SUCCESS}, nil
}

func (s *server) Result(context context.Context, empty *GoAuctionSystem.Empty) (*GoAuctionSystem.Outcome, error) {
	return &GoAuctionSystem.Outcome{Amount: int32(bidVal), Over: false}, nil
}

func StartServer() {
	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", id+5000))
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
	args := os.Args[1:] // args: <port number>
	pid, _ := strconv.ParseInt(args[0], 10, 32)
	id = int32(pid)

	StartServer()
}
