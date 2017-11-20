package service_test

import (
	"strings"
	"testing"

	"github.com/choefele/todo-backend-go/service"
)

func TestValidationCreateTitleLength(t *testing.T) {
	todoService := service.NewValidation(service.NewMemoryService())

	todoForm := service.TodoForm{
		Title: strings.Repeat("#", 257),
	}
	_, err := todoService.Create(nil, todoForm)

	if err == nil {
		t.Error("Expected validation error")
	}
}
