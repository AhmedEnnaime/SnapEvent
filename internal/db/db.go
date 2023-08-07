package db

import (
	"fmt"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(config *configs.Config) *sqlx.DB {

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
