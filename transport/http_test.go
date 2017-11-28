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

	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	server.todos(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return %v but was %v", http.StatusOK, w.Code)
	}

	expected := []byte("[]\n")
	if !bytes.Equal(w.Body.Bytes(), expected) {
		t.Errorf("Body should be `%s` but was `%s`", expected, w.Body.Bytes())
	}
}
