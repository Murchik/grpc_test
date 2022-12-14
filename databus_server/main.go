package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "grpc_test/murchik/test/databus"

	"google.golang.org/grpc"
)

var (
	action       string
	validActions = [...]string{"mul", "div", "add", "sub"}
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
		if in.Rhs == 0 {
			return &pb.SendResponse{Result: r}, errors.New("division by zero")
		}
		r = in.Lhs / in.Rhs
	case "add":
		r = in.Lhs + in.Rhs
	case "sub":
		r = in.Lhs - in.Rhs
	}
	return &pb.SendResponse{Result: r}, nil
}

func validateAction(act string) (string, error) {
	for _, action := range validActions {
		if act == action {
			return act, nil
		}
	}
	return "", errors.New("invalid action")
}

func main() {
	flag.Parse()

	// Parse port from args
	port, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatalf("failed to parse port: %v", err)
	}

	// Parse action from args
	a, err := validateAction(flag.Arg(1))
	if err != nil {
		log.Fatalf("failed to parse action: %v", err)
	}
	action = a

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
