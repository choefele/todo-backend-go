package main

import (
	"fmt"

	"github.com/choefele/todo-backend-go/repository"
	"github.com/choefele/todo-backend-go/service"
	"github.com/choefele/todo-backend-go/transport"
)

const httpPort = 8080
const grpcPort = 8888

func main() {
	// todoService := repository.NewMemoryService()
	todoService := repository.NewMongoService()
	todoService = service.NewValidation(todoService)

	go func() {
		fmt.Printf("HTTP server listening on port %v\n", httpPort)
		transport.NewHTTPServer(todoService).ListenAndServe("/api", httpPort)
	}()

	fmt.Printf("gRPC server listening on port %v\n", grpcPort)
	transport.NewGRPCServer(todoService).ListenAndServe(grpcPort)
}
