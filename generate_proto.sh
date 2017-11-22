#!/bin/bash

protoc -I api/ api/grpc.proto --go_out=plugins=grpc:api