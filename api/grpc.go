package api

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

func (g *GRPCServer) Todos(context context.Context, _ *TodoRequest) (*TodoResponse, error) {
	todos, err := g.service.Todos(context)
	if err != nil {
		return nil, err
	}

	todoGRPC := []*Todo{}
	for _, t := range todos {
		todoGRPC = append(todoGRPC, &Todo{
			Id:    t.ID,
			Title: t.Title,
		})
	}
	todoResponse := TodoResponse{
		Todos: todoGRPC,
	}
	return &todoResponse, nil
}
