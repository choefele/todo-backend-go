// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package transport is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	TodoRequest
	TodoResponse
	Todo
*/
package transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TodoRequest struct {
}

func (m *TodoRequest) Reset()                    { *m = TodoRequest{} }
func (m *TodoRequest) String() string            { return proto.CompactTextString(m) }
func (*TodoRequest) ProtoMessage()               {}
func (*TodoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TodoResponse struct {
	Todos []*Todo `protobuf:"bytes,1,rep,name=todos" json:"todos,omitempty"`
}

func (m *TodoResponse) Reset()                    { *m = TodoResponse{} }
func (m *TodoResponse) String() string            { return proto.CompactTextString(m) }
func (*TodoResponse) ProtoMessage()               {}
func (*TodoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TodoResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type Todo struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Todo) Reset()                    { *m = Todo{} }
func (m *Todo) String() string            { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()               {}
func (*Todo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*TodoRequest)(nil), "transport.TodoRequest")
	proto.RegisterType((*TodoResponse)(nil), "transport.TodoResponse")
	proto.RegisterType((*Todo)(nil), "transport.Todo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TodoService service

type TodoServiceClient interface {
	Todos(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Todos(ctx context.Context, in *TodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := grpc.Invoke(ctx, "/transport.TodoService/Todos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoService service

type TodoServiceServer interface {
	Todos(context.Context, *TodoRequest) (*TodoResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Todos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Todos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.TodoService/Todos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Todos(ctx, req.(*TodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Todos",
			Handler:    _TodoService_Todos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f,
	0x2a, 0x51, 0xe2, 0xe5, 0xe2, 0x0e, 0xc9, 0x4f, 0xc9, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x51, 0x32, 0xe5, 0xe2, 0x81, 0x70, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x54, 0xb9, 0x58,
	0x4b, 0xf2, 0x53, 0xf2, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xf8, 0xf5, 0xe0, 0x3a,
	0xf5, 0xc0, 0xea, 0x20, 0xb2, 0x4a, 0x3a, 0x5c, 0x2c, 0x20, 0xae, 0x10, 0x1f, 0x17, 0x53, 0x66,
	0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x08, 0x17, 0x6b, 0x49,
	0x66, 0x49, 0x4e, 0xaa, 0x04, 0x13, 0x58, 0x08, 0xc2, 0x31, 0x72, 0x87, 0xd8, 0x19, 0x9c, 0x5a,
	0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc1, 0xc5, 0x0a, 0xe2, 0x16, 0x0b, 0x89, 0xa1, 0x9b, 0x0e,
	0x71, 0x94, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x75, 0x49, 0x6c, 0x60, 0xef, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x8a, 0x37, 0x60, 0xdc, 0x00, 0x00, 0x00,
}