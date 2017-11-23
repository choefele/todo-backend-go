#!/bin/bash

// https://github.com/grpc-ecosystem/polyglot
echo '{}' | \
    java -jar ~/Downloads/polyglot.jar \
    --command=call \
    --full_method=api.TodoService/Todos \
    --endpoint=localhost:8888 \
    --use_reflection=false \
    --proto_discovery_root=./api
