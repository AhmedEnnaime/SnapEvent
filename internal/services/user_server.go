package services

import (
	"context"
	"fmt"

	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/AhmedEnnaime/SnapEvent/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	h *Handler
}

func NewUserServer(h *Handler) *UserServer {
	return &UserServer{
		h: h,
	}
}

func (server *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	server.h.logger.Info().Interface("req", req).Msg("Create User")

	user := models.User{
		Name:     req.GetName(),
		Birthday: req.GetBirthday(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Gender:   req.GetGender(),
	}

	err := user.Validate()
	if err != nil {
		msg := err.Error()
		err = fmt.Errorf("validation error: %w", err)
		server.h.logger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	err = user.HashPassword()
	if err != nil {
		msg := "internal server error"
		err := fmt.Errorf("failed to hash password")
		server.h.logger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.Aborted, err.Error())
	}

	err = server.h.us.Create(&user)
	if err != nil {
		msg := "internal server error"
		err := fmt.Errorf("failed to create user. %w", err)
		server.h.logger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.Canceled, msg)
	}

	// token, err :=

	res := &pb.UserResponse{
		User: &pb.User{
			Id:       uint32(user.ID),
			Name:     user.Name,
			Birthday: user.Birthday,
			Email:    user.Email,
			Password: user.Password,
			Gender:   string(user.Gender),
		},
	}
	return res, nil
}

func (server *UserServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	server.h.logger.Info().Interface("req", req).Msg("Get all users")

	users, err := server.h.us.GetAllUsers()
	if err != nil {
		server.h.logger.Error().Err(err).Msg("failed to get users from database")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	pbUsers := make([]*pb.User, 0, len(users))
	for _, user := range users {
		pbUser := &pb.User{
			Id:       uint32(user.ID),
			Name:     user.Name,
			Birthday: user.Birthday,
			Email:    user.Email,
			Password: user.Password,
			Gender:   user.Gender,
			// Omitting other fields for brevity
		}
		pbUsers = append(pbUsers, pbUser)
	}

	res := &pb.GetUsersResponse{
		Users: pbUsers,
	}

	return res, nil

}

// func (server *UserServer) GetUserByID(ctx context.Context, req *pb.GetUserId) (*pb.UserResponse, error) {

// }
