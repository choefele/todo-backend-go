#!/bin/bash

# https://github.com/grpc-ecosystem/polyglot
echo "{'title': 'title'}" | \
    java -jar ~/Downloads/polyglot.jar \
    --command=call \
    --full_method=transport.TodoService/Create \
    --endpoint=localhost:8888 \
    --use_reflection=false \
    --proto_discovery_root=./transport

echo '{}' | \
    java -jar ~/Downloads/polyglot.jar \
    --command=call \
    --full_method=transport.TodoService/Todos \
    --endpoint=localhost:8888 \
    --use_reflection=false \
    --proto_discovery_root=./transport
