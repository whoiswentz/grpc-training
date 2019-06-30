package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type GreetServer struct {
	Network string
	Address string
}

func NewGreetServer(network string, address string) *GreetServer {
	return &GreetServer{Network: network, Address: address}
}

func (*GreetServer) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	response := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}
	return response, nil
}

func (*GreetServer) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
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

func (greetServer *GreetServer) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")

	var result string
	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}

		result = "Hello, " + request.GetGreeting().GetFirstName() + " !"
	}
}

func (*GreetServer) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with a streaming request")

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
			return err
		}

		firstName := request.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		if err := stream.Send(&greetpb.GreetEveryoneResponse{Result: result}); err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
	return nil
}

func main() {
	greetServer := NewGreetServer("tcp", "0.0.0.0:50051")

	listener, err := net.Listen(greetServer.Network, greetServer.Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, greetServer)

	log.Println("Serving in 0.0.0.0:5000")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
