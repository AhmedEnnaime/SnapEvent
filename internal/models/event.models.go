package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type STATUS string

const (
    Open STATUS = "Open"
    Closed STATUS = "Closed"
    Invitation STATUS = "Invitation"
)

type Event struct {
	gorm.Model
	EventDate time.Time `json:"event_date" gorm:"not null"`
	Time string `json:"time" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	Location string `json:"location" gorm:"not null"`
	Poster string `json:"poster"`
	Status STATUS `json:"status" gorm:"not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User User `gorm:"foreignkey:UserId"`
	Users []User `gorm:"many2many:user_events;"`

}

func (e Event) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(
			&e.EventDate,
			validation.Required,
			validation.Date("2006-01-02"),
		),
		validation.Field(
			&e.Time,
			validation.Required,
		),
		validation.Field(
			&e.Description,
			validation.Required,
		),
		validation.Field(
			&e.City,
			validation.Required,
		),
		validation.Field(
			&e.Location,
			validation.Required,
		),
		validation.Field(
			&e.Status,
			validation.Required,
			validation.In("Open", "Closed", "Invitation"),
		),
		validation.Field(
			&e.UserID,
			validation.Required,
		),
	
	)
	
}