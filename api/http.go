package api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/choefele/todo-backend-go/service"
)

type HTTPServer struct {
	service service.TodoService
}

func NewHTTPServer(s service.TodoService) *HTTPServer {
	return &HTTPServer{
		service: s,
	}
}

func (h *HTTPServer) ListenAndServe(root string, port int) {
	mux := http.NewServeMux()

	mux.HandleFunc(root+"/todos", h.todos)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

func (h *HTTPServer) todos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.Todos(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	body := fmt.Sprintf("Todos: %v", todos)
	io.WriteString(w, body)
}
