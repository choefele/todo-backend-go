package service

import "context"

type Todo struct {
	ID    string
	Title string
}

type TodoForm struct {
	Title string
}

type TodoService interface {
	Create(context context.Context, form TodoForm) (Todo, error)
	Todo(context context.Context, id string) (Todo, error)
	Todos(context context.Context) ([]Todo, error)
	//edit(context context.Context, todo Todo) Todo
	//tick(context context.Context, todo Todo, off bool) Todo
	//move()
}
