package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/AhmedEnnaime/SnapEvent/internal/services"
	"github.com/AhmedEnnaime/SnapEvent/internal/store"
	"github.com/AhmedEnnaime/SnapEvent/pb"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

var (
	DB *gorm.DB
	h  *services.Handler
)

func init() {

	var bug error

	DB, bug = db.New()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if DB == nil {
		log.Fatalf("Failed to connect to the database: %v", bug) // Print detailed error message
	}

	if err := db.AutoMigrate(DB); err != nil {
		log.Fatalf("Failed to perform database migrations: %v", err)
	}

	if err := db.Seed(DB); err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}

	userStore := store.NewUserStore(DB)
	eventStore := store.NewEventStore(DB)

	fmt.Println("Connected to postgres successfully")

	h = services.New(&logger, userStore, eventStore)
}

func RunGrpcServer(ctx context.Context) error {

	// server_addr := os.Getenv("GRPC_SERVER_ADDRESS")
	// if server_addr == "" {
	// 	log.Fatal("$GRPC_SERVER_ADDRESS is not set")
	// }

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	userServer := services.NewUserServer(h)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return grpcServer.Serve(listener)

}

func main() {
	ctx := context.Background()
	if err := RunGrpcServer(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
