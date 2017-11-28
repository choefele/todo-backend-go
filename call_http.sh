#!/bin/bash

curl -X POST -d "{\"title\": \"title\"}" localhost:8080/api/todos
curl localhost:8080/api/todos