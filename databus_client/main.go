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

	// TODO: cmd arg1 validation
	addr := flag.Arg(0)

	// cmd arg2 validation
	lhs, err := strconv.ParseFloat(flag.Arg(1), 32)
	if err != nil {
		log.Fatalf("Invalid argument at place 2: %v", err)
	}

	// cmd arg3 validation
	rhs, err := strconv.ParseFloat(flag.Arg(2), 32)
	if err != nil {
		log.Fatalf("Invalid argument at place 3: %v", err)
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
