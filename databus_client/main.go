package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc_test/murchik/test/databus"
)

func main() {
	flag.Parse()

	// Check if number of args is valid
	if len(flag.Args()) != 3 {
		log.Fatalf("Invalid number of arguments")
	}

	addr := flag.Arg(0)

	// Parse 1st param
	lhs, err := strconv.ParseFloat(flag.Arg(1), 32)
	if err != nil {
		log.Fatalf("Invalid lhs argument: %v", err)
	}

	// Parse 2nd param
	rhs, err := strconv.ParseFloat(flag.Arg(2), 32)
	if err != nil {
		log.Fatalf("Invalid rhs argument: %v", err)
	}

	// Establish connection
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDatabusServiceClient(conn)

	// Send a request and wait for a response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ProcessRequest(ctx, &pb.SendRequest{Lhs: float32(lhs), Rhs: float32(rhs)})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	// Print result
	log.Printf("Result: %f", r.GetResult())
}
