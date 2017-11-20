package service

import (
	"context"
	"errors"
	"unicode/utf8"
)

type validation struct {
	TodoService
}

// NewValidation creates a new TodoService that forwards validated input data to the
// next service implementation or returns an error.
func NewValidation(next TodoService) TodoService {
	return TodoService(&validation{
		next,
	})
}

func (v *validation) Create(context context.Context, form TodoForm) (Todo, error) {
	if form.Title == "" {
		return Todo{}, errors.New("Invalid title (title empty string)")
	}
	if utf8.RuneCountInString(form.Title) > 256 {
		return Todo{}, errors.New("Invalid title (title length > 256)")
	}

	return v.TodoService.Create(context, form)
}
