package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/unary/greet/greetpb"
	"log"
	"net"
	"strconv"
	"time"
)

type GreetServer struct {
	Network string
	Address string
}

func (s *GreetServer) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	response := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}
	return response, nil
}

func (s *GreetServer) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", request)

	firstName := request.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		response := &greetpb.GreetManyTimesResponse{Result: result}
		if err := stream.Send(response); err != nil {
			log.Fatalf("Error while streaming: %v", response)
			return err
		}
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	greetServer := GreetServer{
		Network: "tcp",
		Address: "0.0.0.0:50051",
	}

	listener, err := net.Listen(greetServer.Network, greetServer.Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, &greetServer)

	log.Println("Serving in 0.0.0.0:5000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
