package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-training/blog/blogpb"
	"io"
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

	// CREATE
	blog := &blogpb.Blog{
		AuthorId: "Vinicios Wentz",
		Title:    "gRPC Awesomeness",
		Content:  "A lot of text about gRPC",
	}
	createdBlog := createBlog(blog, client)
	log.Printf("Created Blof: %v\n", createdBlog)

	// READ
	readedBlog := readBlog(createdBlog, client)
	log.Printf("Readed Blog: %v\n", readedBlog)

	// UPDATE
	updatedBlog := &blogpb.Blog{
		Id:       readedBlog.GetBlog().GetId(),
		AuthorId: "Vinicios Henrique Wentz",
		Title:    "gRPC Awesomeness 2019 UPDATED",
		Content:  "A lot of text about gRPC in 2019",
	}
	newUpdatedBlog := updateBlog(updatedBlog, client)
	log.Printf("Updated Blog: %v\n", newUpdatedBlog)

	// DELETE
	deleteRequest := &blogpb.DeleteBlogRequest{BlogId: newUpdatedBlog.GetBlog().GetId()}
	deleteResult := deleteBlog(deleteRequest, client)
	log.Printf("Blog was deleted: %v\n", deleteResult)

	// List blogs using server streaming
	listBlogs(client)
}

func createBlog(blog *blogpb.Blog, client blogpb.BlogServiceClient) *blogpb.CreateBlogResponse {
	request := &blogpb.CreateBlogRequest{Blog: blog}
	response, err := client.CreateBlog(context.Background(), request)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return response
}

func readBlog(response *blogpb.CreateBlogResponse, client blogpb.BlogServiceClient) *blogpb.ReadBlogResponse {
	blogId := response.GetBlog().GetId()

	readBLogRequest := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogResponse, err := client.ReadBlog(context.Background(), readBLogRequest)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return readBlogResponse
}

func updateBlog(blogToUpdate *blogpb.Blog, client blogpb.BlogServiceClient) *blogpb.UpdateBlogResponse {
	request := &blogpb.UpdateBlogRequest{Blog: blogToUpdate}
	updateResponse, err := client.UpdateBlog(context.Background(), request)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return updateResponse
}

func deleteBlog(blogToDelete *blogpb.DeleteBlogRequest, client blogpb.BlogServiceClient) *blogpb.DeleteBlogResponse {
	deleteResult, err := client.DeleteBlog(context.Background(), blogToDelete)
	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
	}
	return deleteResult
}

func listBlogs(client blogpb.BlogServiceClient) {
	stream, err := client.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling BlogList RPC: %v", err)
	}

	for {
		result, err := stream.Recv()

		if err != io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v\n", err)
		}

		log.Println(result.GetBlog())
	}
}
