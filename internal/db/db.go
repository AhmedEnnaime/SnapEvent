package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AhmedEnnaime/SnapEvent/internal/configs"
	"github.com/AhmedEnnaime/SnapEvent/internal/models"
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
func NewTestDbB() (*gorm.DB, error)  {
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
		txdb.Register("txdb","postgres", s)
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
