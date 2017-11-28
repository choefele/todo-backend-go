package transport

import (
	"encoding/json"
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

type todo struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func (h *HTTPServer) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	form := service.TodoForm{
		Title: "title",
	}
	t, err := h.service.Create(r.Context(), form)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo{
		ID:    t.ID,
		Title: t.Title,
	})
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

	todosHTTP := []todo{}
	for _, t := range todos {
		todosHTTP = append(todosHTTP, todo{
			ID:    t.ID,
			Title: t.Title,
		})
	}

	json.NewEncoder(w).Encode(todosHTTP)
}
