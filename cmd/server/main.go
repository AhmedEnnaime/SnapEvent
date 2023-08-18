package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/AhmedEnnaime/SnapEvent/internal/services"
	"github.com/AhmedEnnaime/SnapEvent/pb"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

var (
	DB *gorm.DB
	h  *services.Handler
)

func init() {

	var bug error

	DB, bug = db.New()

	if DB == nil {
		log.Fatalf("Failed to connect to the database: %v", bug) // Print detailed error message
	}

	if err := db.AutoMigrate(DB); err != nil {
		log.Fatalf("Failed to perform database migrations: %v", err)
	}

	if err := db.Seed(DB); err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}

	fmt.Println("Connected to postgres successfully")
}

func main() {

	server_addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if server_addr == "" {
		log.Fatal("$GRPC_SERVER_ADDRESS is not set")
	}

	log.Printf("Listening on address %s", server_addr)

	userServer := services.NewUserServer(h)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

	listener, err := net.Listen("tcp", server_addr)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
