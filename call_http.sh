#!/bin/bash

curl -X POST -d "{\"test\": \"that\"}" localhost:8080/api/todos
curl localhost:8080/api/todos