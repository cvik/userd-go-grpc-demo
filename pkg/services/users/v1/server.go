package usersv1

import (
	"context"
	"fmt"
	"net"

	"github.com/cvik/userd-go-grpc-demo/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server ----------

type Server struct {
	UnimplementedUsersServer
	store store.Store
}

func NewServer(store store.Store) *Server {
	return &Server{store: store}
}

func (s *Server) Run(port int) error {
	server := grpc.NewServer()

	RegisterUsersServer(server, s)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	return server.Serve(lis)
}

// impl UserServer interface ----------

func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	id, err := s.store.CreateUser(ctx, in.GetName(), in.GetEmail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &CreateUserResponse{Id: id}, nil
}

func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	user, err := s.store.GetUser(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &GetUserResponse{User: FromInner(*user)}, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *DeleteUserRequest) (*DeleteUserResponse, error) {
	if err := s.store.DeleteUser(ctx, in.GetId()); err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &DeleteUserResponse{}, nil
}

func (s *Server) ListUsers(in *ListUsersRequest, server Users_ListUsersServer) error {
	users, err := s.store.ListUsers(server.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, user := range users {
		if err := server.Send(&ListUsersResponse{User: FromInner(user)}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

// Utility ----------

func FromInner(u store.User) *User {
	return &User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
