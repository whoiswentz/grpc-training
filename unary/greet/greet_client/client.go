package main

import (
	"context"
	"google.golang.org/grpc"
	greetpb2 "grpc-training/unary/greet/greetpb"
	"log"
)

const target = "localhost:50051"

func main() {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := greetpb2.NewGreetServiceClient(conn)

	doUnary(client)
}

func doUnary(client greetpb2.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC...")

	request := &greetpb2.GreetRequest{
		Greeting: &greetpb2.Greeting{
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
