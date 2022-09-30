package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "grpc_test/murchik/test/databus"

	"google.golang.org/grpc"
)

var (
	port   int
	action string
)

type server struct {
	pb.UnimplementedDatabusServiceServer
}

func (s *server) ProcessRequest(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {
	log.Printf("Received lhs: %v, rhs: %v", in.Lhs, in.Rhs)
	var r float32
	switch action {
	case "mul":
		r = in.Lhs * in.Rhs
	case "div":
		r = in.Lhs / in.Rhs
	case "add":
		r = in.Lhs + in.Rhs
	case "sub":
		r = in.Lhs - in.Rhs
	}
	return &pb.SendResponse{Result: r}, nil
}

func main() {
	flag.Parse()

	// cmd arg1 validation
	if p, err := strconv.Atoi(flag.Arg(0)); err == nil {
		port = p
		log.Printf("i=%d, type: %T\n", port, port)
	}

	// TODO: cmd arg2 validation
	action = flag.Arg(1)

	// Open TCP connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start gRPC server
	s := grpc.NewServer()
	pb.RegisterDatabusServiceServer(s, &server{})
	log.Printf("The server is running. Listening on port %v. Selected action: %v.", port, action)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
