package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/unary/calculator/calculatorpb"
	"log"
	"net"
)

const (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type CalculatorServer struct{}

func (*CalculatorServer) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with: %v\n", req)
	result := req.GetValues().GetNumber1() + req.GetValues().GetNumber2()
	response := &calculatorpb.CalculatorResponse{Result: result}
	return response, nil
}

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(server, &CalculatorServer{})

	log.Println("Serving in 0.0.0.0:5000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}