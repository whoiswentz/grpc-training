package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-training/greet/greetpb"
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
	fmt.Printf("Created cliente: %v", client)
}
