package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
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


func main()  {

	configPath := os.Getenv("CONFIG_PATH")
    if configPath == "" {
        log.Fatal("CONFIG_PATH environment variable is not set")
    }

	_, err := configs.LoadConfig(configPath)

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer DB.Close()

	log.Printf("Everything OK")

}