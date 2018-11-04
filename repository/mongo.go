package repository

import (
	"context"

	mgo "gopkg.in/mgo.v2"

	"github.com/choefele/todo-backend-go/service"
)

type mongoService struct {
	session *mgo.Session
}

// NewMongoService creates a new TodoService that connects to a Mongo DB.
func NewMongoService(connection string) (service.TodoService, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	service := service.TodoService(&mongoService{
		session: s,
	})
	return service, nil
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
