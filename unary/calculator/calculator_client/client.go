package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/unary/calculator/calculatorpb"
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

	doCalculation(client)
}

func doCalculation(client calculatorpb.CalculatorServiceClient) {
	log.Println("Starting calculation call...")

	request := &calculatorpb.CalculatorRequest{
		Values: &calculatorpb.Calculator{
			Number1: 25,
			Number2: 25,
		},
	}

	response, err := client.Calculate(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", response.GetResult())
}
