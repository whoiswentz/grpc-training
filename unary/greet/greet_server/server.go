package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	greetpb2 "grpc-training/unary/greet/greetpb"
	"log"
	"net"
)

const (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type Server struct{}

func (*Server) Greet(ctx context.Context, req *greetpb2.GreetRequest) (*greetpb2.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	response := &greetpb2.GreetResponse{
		Result: "Hello " + firstName,
	}
	return response, nil
}

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	greetpb2.RegisterGreetServiceServer(server, &Server{})

	fmt.Println("Serving in 0.0.0.0:5000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
