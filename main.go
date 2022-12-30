package main

import (
	GoAuctionSystem "github.com/Hw5_GoAuctionSystem/proto"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:] // args: <client ID> <server Count>
	aid, _ := strconv.ParseInt(args[0], 10, 32)
	sc, _ := strconv.ParseInt(args[1], 10, 32)
	br, _ := strconv.ParseInt(args[2], 10, 32)
	cid = int32(aid)
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
