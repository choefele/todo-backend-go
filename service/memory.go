package service

import (
	"context"
	"fmt"
	"strconv"
)

type memoryService struct {
	id    int
	list  []*Todo
	index map[string]*Todo
}

// NewMemoryService creates a new TodoService that keeps its data in memory.
func NewMemoryService() TodoService {
	return TodoService(&memoryService{
		index: make(map[string]*Todo),
	})
}

func (m *memoryService) Create(_ context.Context, form TodoForm) (Todo, error) {
	todo := Todo{
		ID:    strconv.Itoa(m.id),
		Title: form.Title,
	}
	m.list = append(m.list, &todo)
	m.index[todo.ID] = &todo
	m.id++
	return todo, nil
}

func (m *memoryService) Todo(_ context.Context, id string) (Todo, error) {
	todo, ok := m.index[id]
	if !ok {
		return Todo{}, fmt.Errorf("Invalid index `%v`", id)
	}

	return *todo, nil
}

func (m *memoryService) Todos(_ context.Context) ([]Todo, error) {
	res := []Todo{}
	for _, todo := range m.list {
		res = append(res, *todo)
	}

	return res, nil
}
