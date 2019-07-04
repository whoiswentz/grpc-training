package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-training/blog/blogpb"
	"grpc-training/blog/model"
)

type BlogServer struct {
	Network string
	Address string
}

func (BlogServer) CreateBlog(context context.Context, request *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
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
