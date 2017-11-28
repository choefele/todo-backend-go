package transport

import (
	"fmt"
	"log"
	"net"

	"github.com/choefele/todo-backend-go/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	service service.TodoService
}

func NewGRPCServer(s service.TodoService) *GRPCServer {
	return &GRPCServer{
		service: s,
	}
}

func (g *GRPCServer) ListenAndServe(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterTodoServiceServer(grpcServer, g)
	grpcServer.Serve(lis)
}

func (g *GRPCServer) Create(ctx context.Context, in *CreateRequest) (*CreateResponse, error) {
	form := service.TodoForm{
		Title: in.Title,
	}
	todo, err := g.service.Create(ctx, form)
	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		Todo: &Todo{
			Id:    todo.ID,
			Title: todo.Title,
		},
	}, nil
}

func (g *GRPCServer) Todos(ctx context.Context, _ *TodoRequest) (*TodoResponse, error) {
	todos, err := g.service.Todos(ctx)
	if err != nil {
		return nil, err
	}

	todosGRPC := []*Todo{}
	for _, t := range todos {
		todosGRPC = append(todosGRPC, &Todo{
			Id:    t.ID,
			Title: t.Title,
		})
	}
	todoResponse := TodoResponse{
		Todos: todosGRPC,
	}
	return &todoResponse, nil
}
