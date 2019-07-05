package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/blog/blogpb"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Blog Client Starting")

	options := grpc.WithInsecure()

	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer clientConn.Close()

	client := blogpb.NewBlogServiceClient(clientConn)

	createdBlog := CreateBlog(client)
	log.Println(createdBlog)

	readedBlog := ReadBlog(createdBlog, client)
	log.Println(readedBlog)
}

func CreateBlog(client blogpb.BlogServiceClient) *blogpb.CreateBlogResponse {
	blog := &blogpb.Blog{
		AuthorId: "Vinicios Wentz",
		Title:    "gRPC Awesomeness",
		Content:  "A lot of text about gRPC",
	}
	response, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return response
}

func ReadBlog(response *blogpb.CreateBlogResponse, client blogpb.BlogServiceClient) interface{} {
	blogId := response.GetBlog().GetId()

	readBLogRequest := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogResponse, err := client.ReadBlog(context.Background(), readBLogRequest)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return readBlogResponse
}