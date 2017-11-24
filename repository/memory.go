package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/choefele/todo-backend-go/service"
)

type memoryService struct {
	id    int
	list  []*service.Todo
	index map[string]*service.Todo
}

// NewMemoryService creates a new TodoService that keeps its data in memory.
func NewMemoryService() service.TodoService {
	return service.TodoService(&memoryService{
		index: make(map[string]*service.Todo),
	})
}

func (m *memoryService) Create(_ context.Context, form service.TodoForm) (service.Todo, error) {
	todo := service.Todo{
		ID:    strconv.Itoa(m.id),
		Title: form.Title,
	}
	m.list = append(m.list, &todo)
	m.index[todo.ID] = &todo
	m.id++
	return todo, nil
}

func (m *memoryService) Todo(_ context.Context, id string) (service.Todo, error) {
	todo, ok := m.index[id]
	if !ok {
		return service.Todo{}, fmt.Errorf("Invalid index `%v`", id)
	}

	return *todo, nil
}

func (m *memoryService) Todos(_ context.Context) ([]service.Todo, error) {
	res := []service.Todo{}
	for _, todo := range m.list {
		res = append(res, *todo)
	}

	return res, nil
}
