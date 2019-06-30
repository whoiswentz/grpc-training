package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-training/calculator/calculatorpb"
	"io"
	"log"
	"math"
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
	log.Printf("Received PrimeNumberDecomposition RPC: %v", request)

	number := request.GetNumber()
	divisor := int64(2)

	for number > 2 {
		if number%divisor == 0 {
			if err := stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{PrimeFactor: divisor}); err != nil {
				log.Fatalf("Error while seding data do server: %v", err)
				return err
			}
			number /= divisor
		} else {
			divisor += 1
			log.Printf("Divisor has increase! | Divisor: %v", divisor)
		}
	}
	return nil
}

func (*CalculatorServer) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	log.Println("Received ComputerAverage RPC")

	sum := int32(0)
	count := 0

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += request.GetNumber()
		count += 1
	}

}

func (*CalculatorServer) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	log.Println("Received FindMaximum RPC")

	maximum := int32(0)

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		if number := request.GetNumber(); number > maximum {
			maximum = number
			if err := stream.Send(&calculatorpb.FindMaximumResponse{Result: number}); err != nil {
				log.Fatalf("Error while sending data to client %v", err)
				return err
			}
		}
	}
}

func (*CalculatorServer) SquareRoot(context context.Context, request *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	log.Println("Received SquareRoot RPC")

	if number := request.GetNumber(); number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	} else {
		return &calculatorpb.SquareRootResponse{NumberRoot: math.Sqrt(float64(number))}, nil
	}
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
