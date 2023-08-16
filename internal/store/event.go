package store

import "github.com/jinzhu/gorm"

type EventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{
		db: db,
	}
}
