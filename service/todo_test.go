package service_test

import (
	"reflect"
	"testing"

	"github.com/choefele/todo-backend-go/service"
)

func TestCreate(t *testing.T) {
	todoService := service.TodoService(service.NewMemoryService())

	todoCreated, err := todoService.Create(nil)
	if err != nil {
		t.Error(err)
	}
	if todoCreated.ID == "" {
		t.Errorf("Invalid ID `%v`", todoCreated.ID)
	}

	todoRetrieved, err := todoService.Todo(nil, todoCreated.ID)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(todoCreated, todoRetrieved) {
		t.Errorf("Created todo %v not equal to retrieved todo %v", todoCreated, todoRetrieved)
	}
}

func TestRetrieve(t *testing.T) {
	todoService := service.TodoService(service.NewMemoryService())

	todoCreated, err := todoService.Create(nil)
	todoRetrieved, err := todoService.Todo(nil, todoCreated.ID)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(todoCreated, todoRetrieved) {
		t.Errorf("Created todo %v not equal to retrieved todo %v", todoCreated, todoRetrieved)
	}
}

func TestRetrieveInvalidID(t *testing.T) {
	todoService := service.TodoService(service.NewMemoryService())

	_, err := todoService.Todo(nil, "invalid.ID")
	if err == nil {
		t.Error("Invalid ID should return error")
	}
}
