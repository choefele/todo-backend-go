package repository

import (
	"context"

	"github.com/choefele/todo-backend-go/service"
)

type mongoService struct {
}

// NewMongoService creates a new TodoService that connects to a Mongo DB.
func NewMongoService() service.TodoService {
	return service.TodoService(&mongoService{})
}

func (m *mongoService) Create(_ context.Context, form service.TodoForm) (service.Todo, error) {
	return service.Todo{}, nil
}

func (m *mongoService) Todo(_ context.Context, id string) (service.Todo, error) {
	return service.Todo{}, nil
}

func (m *mongoService) Todos(_ context.Context) ([]service.Todo, error) {
	res := []service.Todo{}
	res = append(res, service.Todo{})

	return res, nil
}
