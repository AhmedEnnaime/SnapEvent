package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

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

	defer DB.Close()

	log.Printf("Everything OK")

}
