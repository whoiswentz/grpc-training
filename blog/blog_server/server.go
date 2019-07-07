package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-training/blog/blogpb"
	"grpc-training/blog/model"
	"log"
)

type BlogServer struct {
	Network string
	Address string
}

func (*BlogServer) CreateBlog(context context.Context, request *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := request.GetBlog()

	blogItem := model.BlogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetContent(),
		Content:  blog.GetContent(),
	}

	response, err := collection.InsertOne(context, blogItem)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%v\n", err))
	}

	objectId, ok := response.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to ObjectID: %v\n", err))
	}

	blogResponse := &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       objectId.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}

	return blogResponse, nil
}

func (*BlogServer) ReadBlog(context context.Context, request *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	log.Println("Reading blog request")

	blogID, err := primitive.ObjectIDFromHex(request.GetBlogId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	blogItem := &model.BlogItem{}
	filter := bson.M{"_id": blogID}

	result := collection.FindOne(context, filter)
	if err := result.Decode(blogItem); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog with specified ID: %v\n", err))
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       blogItem.ID.Hex(),
			AuthorId: blogItem.AuthorID,
			Title:    blogItem.Title,
			Content:  blogItem.Content,
		},
	}, nil
}

func (*BlogServer) UpdateBlog(context context.Context, request *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	log.Println("Update Blog Request")
	blog := request.GetBlog()

	blogID, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	blogItem := &model.BlogItem{}
	filter := bson.M{"_id": blogID}

	result := collection.FindOne(context, filter)
	if err := result.Decode(blogItem); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog with specified ID: %v\n", err))
	}

	blogItem.AuthorID = blog.GetAuthorId()
	blogItem.Title = blog.GetTitle()
	blogItem.Content = blog.GetContent()

	if _, err := collection.ReplaceOne(context, filter, blogItem); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot update object in MongoDB: %v\n", err))
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       blogItem.ID.Hex(),
			AuthorId: blogItem.AuthorID,
			Title:    blogItem.Title,
			Content:  blogItem.Content,
		},
	}, nil
}

// Delete blog RPC
// Receives a context and a DeleteBlogRequest
// Returns a DeleteBlogResponse and an Error
func (*BlogServer) DeleteBlog(context context.Context, request *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	log.Println("Delete Blog Request")

	blogID, err := primitive.ObjectIDFromHex(request.GetBlogId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	filter := bson.M{"_id": blogID}

	deleteResult, err := collection.DeleteOne(context, filter);
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot Delete object in MongoDB: %v\n", err))
	}

	if deleteResult.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot found blog in MongoDB: %v\n", err))
	}

	return &blogpb.DeleteBlogResponse{BlogId: request.GetBlogId()}, nil
}

// ListBlog RPC
// Uses server streaming to send the blogs back to the client
func (*BlogServer) ListBlog(request *blogpb.ListBlogRequest, server blogpb.BlogService_ListBlogServer) error {
	log.Println("List Blog RPC Called - Using Server Streaming")

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknow internal error: %v\n", err))
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		blog := &model.BlogItem{}
		if err := cursor.Decode(blog); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v\n", err))
		}

		listBlogResponse := &blogpb.ListBlogResponse{
			Blog: &blogpb.Blog{
				Id:       blog.ID.Hex(),
				AuthorId: blog.AuthorID,
				Title:    blog.Title,
				Content:  blog.Content,
			},
		}

		if err := server.Send(listBlogResponse); err != nil {
			return status.Errorf(codes.DataLoss, fmt.Sprintf("Error while streaming to client: %v\n", err))
		}

		if err := cursor.Err(); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Unknow internal error: %v\n", err))
		}
	}
	return nil
}
