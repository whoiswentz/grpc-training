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
	//doServerStreaming(client)
	doClientStreaming(client)
}

func doClientStreaming(client calculatorpb.CalculatorServiceClient) {
	log.Println("Starting ComputeAverage Client Streaming RPC...")

	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while oppening stream: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 5, 6}

	for _, number := range numbers {
		log.Printf("Sending number: %v", number)
		_ = stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("The average is: %v", response.GetAverage())
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
