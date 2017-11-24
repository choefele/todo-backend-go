#!/bin/bash

protoc -I transport/ transport/grpc.proto --go_out=plugins=grpc:transport