package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-training/greet/greetpb"
	"log"
	"net"
)

const (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type Server struct{}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	response := &greetpb.GreetResponse{
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
	greetpb.RegisterGreetServiceServer(server, &Server{})

	fmt.Println("Serving in 0.0.0.0:5000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
