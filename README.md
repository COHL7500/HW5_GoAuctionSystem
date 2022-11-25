# HW5_GoAuctionSystem
Homework 5 - Distributed Systems - Distributed Auction System

# How to start our program
## Step 1.
Firstly, open op x amount servers and y amount of clients, we recommend 3 servers and 3 clients where each client bid in a total of 2 rounds.

## Step 2.
Locate server.go and client.go from the location you have downloaded it into.

## Step 3.
Enter the following code:
```
go run client.go <client_number> <total_servers> <auction_rounds>
go run server.go <server_number>
```

## Step 3.1
The below is how you would start our program using 6 consoles:

```
// Console number 1
go run server.go 0

// Console number 2
go run server.go 1

// Console number 3
go run server.go 2

// Console number 4
go run client.go 0 3 2

// Console number 5
go run client.go 1 3 2

// Console number 6
go run client.go 2 3 2

```

# Step 3.2
From here when everything is up and running, a crash can be simulated by pressing CTRL+C in console number 4 for the server with id 0

## Final step
You should be done, the result can be compared with the text from our log.txt file uploaded with our .pdf documentation
