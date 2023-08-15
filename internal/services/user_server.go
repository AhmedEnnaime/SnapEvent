package services

import (
	"context"

	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/AhmedEnnaime/SnapEvent/pb"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (server *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		Name:     req.GetName(),
		Birthday: req.GetBirthday().AsTime(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Gender:   models.GENDER(req.GetGender().String()),
	}

	err := user.Validate()
	if err != nil {

	}
}
