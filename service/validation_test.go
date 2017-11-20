package service_test

import (
	"strings"
	"testing"

	"github.com/choefele/todo-backend-go/service"
)

func TestValidationCreateTitle(t *testing.T) {
	todoService := service.NewValidation(service.NewMemoryService())

	var testCases = []struct {
		name   string
		title  string
		result bool
	}{
		{"Length == 256", strings.Repeat("#", 256), true},
		{"Length > 256", strings.Repeat("#", 257), false},
		{"Empty string", "", false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := todoService.Create(nil, service.TodoForm{
				Title: testCase.title,
			})

			result := err == nil
			if result != testCase.result {
				t.Errorf("Should be %t but was %t", testCase.result, result)
			}
		})
	}
}
