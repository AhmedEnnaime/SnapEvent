package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserClient struct {
	service pb.UserServiceClient
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	service := pb.NewUserServiceClient(cc)
	return &UserClient{service: service}
}

func (userClient *UserClient) SignUp(user *pb.User) {
	req := &pb.CreateUserRequest{
		Name:     user.Name,
		Birthday: user.Birthday,
		Email:    user.Email,
		Password: user.Password,
		Gender:   user.Gender,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := userClient.service.CreateUser(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("user already exists")
		} else {
			log.Fatal("cannot create user: ", err)
		}
		return
	}

	fmt.Println(res)

}

// func (userClient *UserClient) GetUsers(ctx context.Context, req *pb.GetUsersRequest, opts ...grpc.CallOption) ([]*pb.User, error) {
// 	res, err := userClient.service.GetUsers(ctx, req, opts...)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get users: %v", err)
// 	}

// 	return res.Users, nil
// }

func (userClient *UserClient) GetUserByID(args *pb.GetUserId) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	res, err := userClient.service.GetUserById(ctx, args)
	if err != nil {
		log.Fatalf("GetUser: %v", err)
	}

	fmt.Println(res)
}
