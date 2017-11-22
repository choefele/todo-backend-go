package main

import (
	"fmt"

	"github.com/choefele/todo-backend-go/api"
	"github.com/choefele/todo-backend-go/service"
)

const port = 8080

func main() {
	todoService := service.NewMemoryService()
	todoService = service.NewValidation(todoService)
	fmt.Printf("HTTP server listening on port %v", port)
	api.NewHTTPServer(todoService).ListenAndServe("/api", port)
}
