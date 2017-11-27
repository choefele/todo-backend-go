package transport_test

import (
	"testing"

	"github.com/choefele/todo-backend-go/repository"
	"github.com/choefele/todo-backend-go/transport"
)

func TestTodos(t *testing.T) {
	todoService := repository.NewMemoryService()
	server := transport.NewGRPCServer(todoService)

	request := transport.TodoRequest{}
	response, err := server.Todos(nil, &request)

	if err != nil {
		t.Error(err)
	}

	if len(response.Todos) != 0 {
		t.Errorf("Expected 0 todos but received %v", len(response.Todos))
	}
}
