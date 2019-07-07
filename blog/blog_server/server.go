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
