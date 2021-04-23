package services

import (
	"context"
	"fmt"

	"github.com/diltheyaislan/grpc-golang/pb/pb"
)

type UserServiceServer interface {
	AddUser(context.Context, *pb.User) (*pb.User, error)
	mustEmbedUnimplementedUserServiceServer()
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(cxt context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
