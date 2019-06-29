package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-training/greet/greetpb"
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

	client := greetpb.NewGreetServiceClient(conn)

	//doUnary(client)
	//doServerStreaming(client)
	doClientStreaming(client)
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Giselle",
				LastName:  "Piasetzki",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Vinicios",
				LastName:  "Wentz",
			},
		},
	}

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, request := range requests {
		stream.Send(request)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	fmt.Println("LongGreet response: ", response.GetResult())
}

func doServerStreaming(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a Server Streaming RPC...")

	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vinicios",
			LastName:  "Wentz",
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the server-stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", message.GetResult())
	}
}

func doUnary(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC...")

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vinicios",
			LastName:  "Wentz",
		},
	}

	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", response.GetResult())
}
