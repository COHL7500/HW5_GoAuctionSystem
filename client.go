// --------------------------- //
// ---------- IMPORT --------- //
// --------------------------- //
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Hw5_GoAuctionSystem/proto"
	"google.golang.org/grpc"
)

// --------------------------- //
// --------- GLOBALS --------- //
// --------------------------- //
var (
	cid         int32
	lamport     = 0
	serverCount = 1
	mu          sync.Mutex
	currentBid  = 0
	servers     = make([]GoAuctionSystem.AuctionSystemClient, serverCount)
)

// --------------------------- //
// ---------- CLIENT --------- //
// --------------------------- //
func HandleMessage(serverId int, bid GoAuctionSystem.BidPost) {
	lamport++
	if bid.Amount == -1 {
		log.Printf("Bidding amount %d...", bid.Amount)
		switch ack, _ := clients[serverId].Bid(context.Background(), &bid); ack.Ack {
		case GoAuctionSystem.Acks_ACK_FAIL:
			log.Printf("Bid failed")
		case GoAuctionSystem.Acks_ACK_SUCCESS:
			log.Printf("Bid success")
		case GoAuctionSystem.Acks_ACK_EXCEPTION:
			log.Printf("Bid exception")
		}
		return
	}
	// result, _ := client.Result(context.Background(),&bid)
}

func GetResult() *GoAuctionSystem.Outcome {
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	result, err := clients[0].Result(timeout, &GoAuctionSystem.Empty{})
	if err != nil {

	}

	return result
}

func FrontEnd() {
	for {
		if result := GetResult(); result.Over {

		}
	}
}

func DialServer(serverId int) {
	// Dial server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:", serverId+5000), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// Setup client
	servers[id] = GoAuctionSystem.NewAuctionSystemClient(conn)
	log.Printf("Client connected to server...")

	// Closes connection
	s := make(chan bool)
	<-s
}

// --------------------------- //
// ---------- SETUP ---------- //
// --------------------------- //
func main() {
	args := os.Args[1:] // args: <client ID>
	id, _ := strconv.ParseInt(args[0], 10, 32)
	cid = int32(id)

	for i := 0; i < serverCount; i++ {
		DialServer(i)
	}

	for i := 0; i < 5; i++ {
		FrontEnd()
	}
}
