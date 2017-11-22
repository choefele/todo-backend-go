package main

import (
	"fmt"

	"github.com/choefele/todo-backend-go/api"
	"github.com/choefele/todo-backend-go/service"
)

const httpPort = 8080
const grpcPort = 8888

func main() {
	todoService := service.NewMemoryService()
	todoService = service.NewValidation(todoService)

	go func() {
		fmt.Printf("HTTP server listening on port %v\n", httpPort)
		api.NewHTTPServer(todoService).ListenAndServe("/api", httpPort)
	}()

	fmt.Printf("gRPC server listening on port %v\n", grpcPort)
	api.NewGRPCServer(todoService).ListenAndServe(grpcPort)
}
