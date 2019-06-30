package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"grpc-training/greet/greetpb"
	"io"
	"log"
	"time"
)

const target = "localhost:50051"

func main() {

	crtFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(crtFile, "")
	if sslErr != nil {
		log.Fatalf("Error while loading CA trust certificates: %v", sslErr)
		return
	}

	opts := grpc.WithTransportCredentials(creds)
	conn, err := grpc.Dial(target, opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	doUnary(client)
	//doServerStreaming(client)
	//doClientStreaming(client)
	//doBiDiStreaming(client)

	//doUnaryWithDeadline(client, 1 * time.Second)
	//doUnaryWithDeadline(client, 6 * time.Second)
}

func doUnaryWithDeadline(client greetpb.GreetServiceClient, timeout time.Duration) {
	log.Println("Starting unary RPC with deadline")

	request := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vinicios",
			LastName:  "Wentz",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := client.GreetWithDeadline(ctx, request)

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Timeout has hit! Deadline exceeded!")
			} else {
				log.Fatalf("Unexpected error: %v\n", statusErr)
			}
		} else {
			log.Fatalf("Error while calling GreetWithDeadline RPC: %v\n", err)
		}
		return
	}
	log.Printf("Response from GreetWithDeadline: %v", response.GetResult())
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
