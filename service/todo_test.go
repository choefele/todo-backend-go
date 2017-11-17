package service_test

import (
	"reflect"
	"testing"

	"github.com/choefele/todo-backend-go/service"
)

func TestCreate(t *testing.T) {
	todoService := service.TodoService(service.NewMemoryService())

	todoCreate, err := todoService.Create(nil)
	if err != nil {
		t.Error(err)
	}

	todoRetrieved, err := todoService.Todo(nil, todoCreate.ID)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(todoCreate, todoRetrieved) {
		t.Errorf("Created todo %v not equal to retrieved todo %v", todoCreate, todoRetrieved)
	}
}

func TestRetrieveInvalidID(t *testing.T) {
	todoService := service.TodoService(service.NewMemoryService())

	_, err := todoService.Todo(nil, "invalid.ID")
	if err == nil {
		t.Error("Invalid ID should return error")
	}
}
