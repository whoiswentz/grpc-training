#!/usr/bin/env bash

protoc stream/greet-many-times/greetmanypb/greet-many-times.proto --go_out=plugins=grpc:.