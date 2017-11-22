package api

import (
	"fmt"
	"log"
	"net"

	"github.com/choefele/todo-backend-go/service"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
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

func (g *GRPCServer) Todos(context.Context, *google_protobuf.Empty) (*TodoResponse, error) {
	return nil, nil
}
