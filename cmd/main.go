package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	ctx    context.Context
	DB     *gorm.DB
)

func init() {
	
	ctx = context.TODO()

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

	server = gin.Default()
}

func main() {
	configPath := os.Getenv("CONFIG_PATH")
    if configPath == "" {
        log.Fatal("CONFIG_PATH environment variable is not set")
    }

	config, err := configs.LoadConfig(configPath)

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer DB.Close()

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello from 9090 port"})
	})

	log.Fatal(server.Run(":" + strconv.Itoa(config.Port)))

}
