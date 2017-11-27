package transport

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choefele/todo-backend-go/repository"
)

func TestTodos(t *testing.T) {
	todoService := repository.NewMemoryService()
	server := NewHTTPServer(todoService)
	handler := http.HandlerFunc(server.todos)

	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return %v but was %v", http.StatusOK, w.Code)
	}

	expected := []byte(`Todos: []`)
	if !bytes.Equal(w.Body.Bytes(), expected) {
		t.Errorf("Body should be `%s` but was %v", w.Body.Bytes(), expected)
	}
}