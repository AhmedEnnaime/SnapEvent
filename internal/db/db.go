package db

import (
	"fmt"
	"log"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	config, err := configs.LoadConfig("../../app.env")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	DB = getConnection(&config)
	return DB
}

func getConnection(config *configs.Config) *sqlx.DB {
	var dbConnectionStr string

	dbConnectionStr = fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.PostgresDb,
		config.PostgresUser,
		config.PostgresPassword,
	)

	db, err := sqlx.Open("postgres", dbConnectionStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxIdleConns(5)

	fmt.Println("Connected to DB")

	return db

}
