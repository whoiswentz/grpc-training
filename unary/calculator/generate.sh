#!/usr/bin/env bash

protoc unary/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.