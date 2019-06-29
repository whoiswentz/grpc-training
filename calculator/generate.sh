#!/usr/bin/env bash

protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.