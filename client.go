// --------------------------- //
// ---------- IMPORT --------- //
// --------------------------- //
package main

import (
    "context"
    "log"
    "os"
    "strconv"

    "github.com/Hw5_GoAuctionSystem/proto"
    "google.golang.org/grpc"
)

// --------------------------- //
// --------- GLOBALS --------- //
// --------------------------- //
var (
    id int32
    lamport int64
    currentBid int32
    chanDone chan bool
    chanOut chan GoAuctionSystem.BidPost
    client GoAuctionSystem.AuctionSystemClient
)

// --------------------------- //
// ---------- CLIENT --------- //
// --------------------------- //
func SetBid(amount int32){
    chanOut <-GoAuctionSystem.BidPost{Id: id, Amount: amount, Lamport: lamport}
}

func GetResult(){
    chanOut <-GoAuctionSystem.BidPost{Id: id, Amount: -1, Lamport: lamport}
}

func HandleMessage(bid GoAuctionSystem.BidPost){
    lamport++
    if bid.Amount == -1 {
        log.Printf("Bidding amount %d...", bid.Amount)
        switch ack, _ := client.Bid(context.Background(),&bid); ack.Ack {
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

func FrontEnd() {
    // TBD
}

func StartClient() {
    // Dial server
    conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }

    // Setup client
    client = GoAuctionSystem.NewAuctionSystemClient(conn)
    log.Printf("Client connected to server...")

    // Closes connection
    for {
        select {
           case b := <-chanOut:
                HandleMessage(b)
           case <-chanDone:
               log.Printf("Client closing connection...")
               conn.Close()
               return
        }
    }
}

// --------------------------- //
// ---------- SETUP ---------- //
// --------------------------- //
func main() {
    args := os.Args[1:] // args: <client ID>
    cid, _ := strconv.ParseInt(args[0],10,32)
    id = int32(cid)
    lamport = 0
    currentBid = 0
    chanDone = make(chan bool)
    chanOut = make(chan GoAuctionSystem.BidPost)

    StartClient()
}
