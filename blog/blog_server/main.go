package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpc-training/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
)

var collection *mongo.Collection

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	const URL = "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(URL)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error while creating MongoDB connection: %v\n", err)
	}

	log.Println("Connection to MongoDB")
	if err := client.Connect(context.Background()); err != nil {
		log.Fatalf("Error while connecting to MongoDB: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	log.Println("Blog Service Started")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	server := grpc.NewServer([]grpc.ServerOption{}...)

	blogpb.RegisterBlogServiceServer(server, &BlogServer{})

	go func(server *grpc.Server) {
		log.Println("Starting the server")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}(server)

	// Wait for Control-C to exit
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)

	// Block until a signal is received
	<-interruptChannel

	log.Println("Stopping the server")
	server.Stop()

	log.Println("Closing the listener")
	_ = listener.Close()

	log.Println("Closing MongoDB connection")
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error while closing MongoDB connhection: %v\n", err)
	}

	log.Println("Program terminated")
}
