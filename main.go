package main

import (
	"context"
	"fmt"
	"github.com/grpc_server/server/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	user.UnimplementedInventoryServer
}

func (s *server) GetUserList(ctx context.Context, in *user.GetUserListRequest) (*user.GetUserListResponse, error) {
	fmt.Printf("request - %+v", in)
	if in.GetName() != "" {
		return &user.GetUserListResponse{
			Users: []*user.User{{
				Name:          "Harish",
				Gender:        "Male",
				Age:           28,
				MaritalStatus: &[]string{"Married"}[0],
			}},
		}, nil
	}
	return &user.GetUserListResponse{
		Users: []*user.User{
			{
				Name:          "Harish",
				Gender:        "Male",
				Age:           28,
				MaritalStatus: &[]string{"Married"}[0],
			},
			{
				Name:          "John",
				Gender:        "Male",
				Age:           27,
				MaritalStatus: &[]string{"Unmarried"}[0],
			},
			{
				Name:          "Swathi",
				Gender:        "Female",
				Age:           27,
				MaritalStatus: &[]string{"Unmarried"}[0],
			},
			{
				Name:          "Kavya",
				Gender:        "Female",
				Age:           26,
				MaritalStatus: &[]string{"Married"}[0],
			},
		},
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	user.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
