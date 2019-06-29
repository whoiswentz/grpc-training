package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/calculator/calculatorpb"
	"log"
	"net"
)

const (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type CalculatorServer struct {
	Network string
	Address string
}

func NewCalculatorServer(network string, address string) *CalculatorServer {
	return &CalculatorServer{Network: network, Address: address}
}

func (*CalculatorServer) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with: %v\n", req)
	result := req.GetValues().GetNumber1() + req.GetValues().GetNumber2()
	response := &calculatorpb.CalculatorResponse{Result: result}
	return response, nil
}

func (*CalculatorServer) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Recaived PrimeNumberDecomposition RPC: %v", request)

	number := request.GetNumber()
	divisor := int64(2)

	for number > 2 {
		if number%divisor == 0 {
			_ = stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number /= divisor
		} else {
			divisor += 1
			log.Printf("Divisor has increase! | Divisor: %v", divisor)
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	calculatorServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(calculatorServer, &CalculatorServer{})

	log.Println("Serving in 0.0.0.0:5000")
	if err := calculatorServer.Serve(listener); err != nil {
		log.Fatalf("Failed to calculatorServer: %v", err)
	}
}
