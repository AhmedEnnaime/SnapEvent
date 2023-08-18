package main

import (
	"context"
	"log"
	"os"

	"github.com/AhmedEnnaime/SnapEvent/pb"
	"github.com/AhmedEnnaime/SnapEvent/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	clientAddr := os.Getenv("CLIENT_ORIGIN")
	if clientAddr == "" {
		log.Fatal("$CLIENT_ORIGIN is not set")
	}

	log.Printf("client address: %s", clientAddr)

	conn, err := grpc.Dial(clientAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	userClient := pb.NewUserServiceClient(conn)

	user := utils.NewUser()
	req := &pb.CreateUserRequest{
		Name:     user.Name,
		Birthday: user.Birthday,
		Email:    user.Email,
		Password: user.Password,
		Gender:   user.Gender,
	}

	res, err := userClient.CreateUser(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("user already exists")
		} else {
			log.Fatal("cannot create user ", err)
		}
		return
	}
	log.Printf("user created %s", res)

}
