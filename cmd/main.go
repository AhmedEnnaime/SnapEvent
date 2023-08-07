package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	ctx    context.Context
	DB     *sqlx.DB
)

func init() {
	config, err := configs.LoadConfig("..")

	if err != nil {
		log.Fatal("Could not load environment variables ", err)
	}

	ctx = context.TODO()

	DB := db.ConnectDB(&config)

	if DB != nil {
		log.Fatal("Failed to connect to the database")
	}

	fmt.Println("Connected to postgres successfully")

	server = gin.Default()
}

func main() {
	config, err := configs.LoadConfig("..")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer DB.Close()

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Hello from PostgreSql"})
	})

	log.Fatal(server.Run(":" + strconv.Itoa(config.Port)))

}
