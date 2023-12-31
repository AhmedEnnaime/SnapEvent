package main

import (
	"context"
	"log"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/pb"
	"github.com/AhmedEnnaime/SnapEvent/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	defer conn.Close()
	log.Println(" connection state ====> ", conn.GetState())
	userClient := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := utils.NewUser()
	req := &pb.CreateUserRequest{
		Name:     user.Name,
		Birthday: user.Birthday,
		Email:    user.Email,
		Password: user.Password,
		Gender:   user.Gender,
	}

	res, err := userClient.CreateUser(ctx, req)

	if err != nil {
		log.Fatalf("Create user failed: %v", err)
	}
	log.Printf("user created %s", res)

}
