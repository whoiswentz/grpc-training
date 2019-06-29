package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/calculator/calculatorpb"
	"io"
	"log"
)

const target = "localhost:50051"

func main() {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)

	//doUnary(client)
	doServerStreaming(client)
}

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	log.Println("Starting PrimeDecomposition RPC...")

	request := &calculatorpb.PrimeNumberDecompositionRequest{Number: 121212121212}

	stream, err := client.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming PrimeDecomposition: %v", err)
		}
		log.Printf("Response from PrimeDecomposition: %v", response.GetPrimeFactor())
	}
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	log.Println("Starting calculation call...")

	request := &calculatorpb.CalculatorRequest{
		Values: &calculatorpb.Calculator{
			Number1: 25,
			Number2: 25,
		},
	}

	response, err := client.Calculate(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Calculation RPC: %v", err)
	}
	log.Printf("Response from Calculation: %v", response.GetResult())
}
