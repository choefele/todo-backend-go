#!/bin/bash

curl -X POST -d "{\"test\": \"that\"}" localhost:8080/api/todo
echo
curl localhost:8080/api/todos
echo