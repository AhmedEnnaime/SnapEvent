package main

import (
	"fmt"
	"log"

	"github.com/AhmedEnnaime/SnapEvent/internal/db"
)


func init() {

	DB, err := db.New()

	if DB == nil {
        log.Fatalf("Failed to connect to the database: %v", err) // Print detailed error message
    }

	if err := db.AutoMigrate(DB); err != nil {
        log.Fatalf("Failed to perform database migrations: %v", err)
    }

	if err := db.Seed(DB); err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}

	fmt.Println("Connected to postgres successfully")
}
