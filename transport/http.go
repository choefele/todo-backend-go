package transport

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
	mux.HandleFunc(root+"/todo", h.create)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

func (h *HTTPServer) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	form := service.TodoForm{
		Title: "title",
	}
	todo, err := h.service.Create(r.Context(), form)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf("Todo: %v", todo)
	io.WriteString(w, body)
}

func (h *HTTPServer) todos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	todos, err := h.service.Todos(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body := fmt.Sprintf("Todos: %v", todos)
	io.WriteString(w, body)
}
