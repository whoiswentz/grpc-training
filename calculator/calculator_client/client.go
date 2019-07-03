package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"grpc-training/calculator/calculatorpb"
	"io"
	"log"
	"time"
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
	//doClientStreaming(client)
	_ = doBiDiStreaming(client)
	//doErrorUnarySquareRoot(client)
}

func doErrorUnarySquareRoot(client calculatorpb.CalculatorServiceClient) {
	log.Println("Starting SquareRoot unary RPC...")

	calculatorResponse := &calculatorpb.SquareRootRequest{Number: 10}

	doErrorCall(client, calculatorResponse)

	calculatorResponse = &calculatorpb.SquareRootRequest{Number: -10}

	doErrorCall(client, calculatorResponse)
}

func doErrorCall(client calculatorpb.CalculatorServiceClient, calculatorResponse *calculatorpb.SquareRootRequest) {
	response, err := client.SquareRoot(context.Background(), calculatorResponse)

	if err != nil {
		responseError, ok := status.FromError(err)
		if ok {
			log.Fatalf("%v: %v\n", responseError.Message(), responseError.Code())
		} else {
			log.Fatalf("Error calling SquareRoot: %v\n", err)
		}
	}
	fmt.Printf("Result of square root of: %v: %v\n", calculatorResponse.GetNumber(), response.GetNumberRoot())
}

func doBiDiStreaming(client calculatorpb.CalculatorServiceClient) error {
	log.Println("Starting FindMaximum Bidirectional Streaming RPC...")

	if stream, err := client.FindMaximum(context.Background()); err != nil {
		log.Fatalf("Error while opening stream an calling FindMaximum: %v\n", err)
		return err
	} else {
		waitChannel := make(chan struct{})

		go func() {
			numbers := []int32{13, 200, 31, 4, 1, 1, 70, 2, 500, 12, 1000}
			for _, number := range numbers {
				stream.Send(&calculatorpb.FindMaximumRequest{Number: number})
				time.Sleep(time.Second)
			}
			stream.CloseSend()
		}()

		go func() {
			for {
				if response, err := stream.Recv(); err == io.EOF {
					break
				} else if err != nil {
					log.Fatalf("Problem while reading the server stream: %v\n", err)
					break
				} else {
					maximum := response.GetResult()
					log.Printf("Received a new maximum: %v\n", maximum)
				}
			}
			close(waitChannel)
		}()

		<-waitChannel
	}
	return nil
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
