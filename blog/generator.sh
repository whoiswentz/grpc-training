#!/usr/bin/env bash

protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.