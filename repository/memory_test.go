package repository_test

import (
	"reflect"
	"testing"

	"github.com/choefele/todo-backend-go/repository"
	"github.com/choefele/todo-backend-go/service"
)

func TestCreate(t *testing.T) {
	todoService := repository.NewMemoryService()

	form := service.TodoForm{"title"}
	todo, err := todoService.Create(nil, form)
	if err != nil {
		t.Error(err)
	}
	if todo.ID == "" {
		t.Errorf("Invalid ID `%v`", todo.ID)
	}
	if todo.Title != form.Title {
		t.Error("Invalid property")
	}
}

func TestTodo(t *testing.T) {
	todoService := repository.NewMemoryService()

	todoCreated, err := todoService.Create(nil, service.TodoForm{
		Title: "title",
	})
	todoRetrieved, err := todoService.Todo(nil, todoCreated.ID)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(todoCreated, todoRetrieved) {
		t.Errorf("Created todo %v not equal to retrieved todo %v", todoCreated, todoRetrieved)
	}
}

func TestTodoInvalidID(t *testing.T) {
	todoService := repository.NewMemoryService()

	_, err := todoService.Todo(nil, "invalid.ID")
	if err == nil {
		t.Error("Invalid ID should return error")
	}
}
func TestTodos(t *testing.T) {
	todoService := repository.NewMemoryService()

	todoService.Create(nil, service.TodoForm{})
	todoService.Create(nil, service.TodoForm{})
	todoService.Create(nil, service.TodoForm{})
	todos, err := todoService.Todos(nil)
	if err != nil {
		t.Error(err)
	}
	if len(todos) != 3 {
		t.Error("Invalid number of todos")
	}
}
