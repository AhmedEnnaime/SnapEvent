package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/AhmedEnnaime/SnapEvent/internal/models"
	"github.com/BurntSushi/toml"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var txdbInitialized bool
var mutex sync.Mutex

func dsn() (string, error) {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}
	configuration, err := configs.LoadConfig(configPath)

	if err != nil {
		return "Cannot load env variables", err
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuration.DbHost, configuration.DbPort, configuration.PostgresUser, configuration.PostgresPassword, configuration.PostgresDb), nil

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
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}
	_, err := configs.LoadConfig(configPath)

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
		&models.UserEvent{},
		&models.Invite{},
	).Error
	if err != nil {
		return err
	}
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

	fmt.Printf("Decoded Users: %#v\n", data.Users)
	fmt.Printf("Decoded Events: %#v\n", data.Events)
	fmt.Printf("Decoded Invites: %#v\n", data.Invites)

	for _, u := range data.Users {
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}

	for _, e := range data.Events {
		fmt.Printf("Creating Event: %+v\n", e)
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
