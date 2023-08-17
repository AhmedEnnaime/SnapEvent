package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/BurntSushi/toml"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var txdbInitialized bool
var mutex sync.Mutex

func dsn() (string, error) {

	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.New("$DB_HOST is not set")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return "", errors.New("$POSTGRES_USER is not set")
	}

	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return "", errors.New("$POSTGRES_PASSWORD is not set")
	}

	name := os.Getenv("POSTGRES_DB")
	if name == "" {
		return "", errors.New("$POSTGRES_DB is not set")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.New("$DB_PORT is not set")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name), nil

}

func New() (*gorm.DB, error) {
	s, err := dsn()
	if err != nil {
		return nil, err
	}

	var d *gorm.DB
	for i := 0; i < 10; i++ {
		d, err = gorm.Open("postgres", s)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)

	}

	if err != nil {
		return nil, err
	}

	d.DB().SetMaxIdleConns(3)
	d.LogMode(false)
	return d, nil

}

// For test purposes => Creating isolated db connection for testing
func NewTestDbB() (*gorm.DB, error) {

	err := godotenv.Load("../../test.env")
	if err != nil {
		return nil, err
	}

	s, err := dsn()

	if err != nil {
		return nil, err
	}

	mutex.Lock()
	if !txdbInitialized {
		_d, err := gorm.Open("postgres", s)
		if err != nil {
			return nil, err
		}
		AutoMigrate(_d)
		txdb.Register("txdb", "postgres", s)
		txdbInitialized = true
	}
	mutex.Unlock()

	c, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		return nil, err
	}

	d, err := gorm.Open("postgres", c)
	if err != nil {
		return nil, err
	}

	d.DB().SetMaxIdleConns(3)
	d.LogMode(false)

	return d, nil
}

func DropTestDB(d *gorm.DB) error {
	d.Close()
	return nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.Invite{},
	).Error
	if err != nil {
		return err
	}

	// Set foreign key constraints
	db.Model(&models.Event{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")
	db.Model(&models.Invite{}).AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")
	db.Model(&models.Invite{}).AddForeignKey("event_id", "events(ID)", "CASCADE", "CASCADE")
	db.Table("user_events").AddForeignKey("user_id", "users(ID)", "CASCADE", "CASCADE")
	db.Table("user_events").AddForeignKey("event_id", "events(ID)", "CASCADE", "CASCADE")

	return nil
}

func Seed(db *gorm.DB) error {
	data := struct {
		Users   []models.User
		Events  []models.Event
		Invites []models.Invite
	}{
		Users:   []models.User{},
		Events:  []models.Event{},
		Invites: []models.Invite{},
	}

	files, err := os.ReadDir("internal/db/seed")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".toml") {
			bs, err := os.ReadFile(filepath.Join("internal/db/seed", file.Name()))
			if err != nil {
				return err
			}
			if _, err := toml.Decode(string(bs), &data); err != nil {
				return err
			}
		}
	}

	for _, u := range data.Users {
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}

	for _, e := range data.Events {
		if err := db.Create(&e).Error; err != nil {
			return err
		}
	}

	for _, i := range data.Invites {
		if err := db.Create(&i).Error; err != nil {
			return err
		}
	}

	return nil
}
