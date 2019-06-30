package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-training/greet/greetpb"
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

	client := greetpb.NewGreetServiceClient(conn)

	//doUnary(client)
	//doServerStreaming(client)
	//doClientStreaming(client)
	doBiDiStreaming(client)
}

func doBiDiStreaming(client greetpb.GreetServiceClient) {
	log.Fatalln("Starting to do a Bidirectional streaming RPC...")

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Giselle",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Vinicios",
			},
		},
	}

	waitChannel := make(chan struct{})

	// Sending
	go func(stream greetpb.GreetService_GreetEveryoneClient, requests []*greetpb.GreetEveryoneRequest) {
		for _, request := range requests {
			if err := stream.Send(request); err != nil {
				log.Fatalf("Error while sending request to server: %v", err)
				return
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error while Closing and Sending: %v", err)
			return
		}
	}(stream, requests)

	// Receiving
	go func(stream greetpb.GreetService_GreetEveryoneClient, waitChannel chan struct{}) {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Receiving: %v\n", response.GetResult())
		}
		close(waitChannel)
	}(stream, waitChannel)

	<-waitChannel
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
