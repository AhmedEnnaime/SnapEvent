package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
) 

type TYPE string

const (
    Attendee TYPE = "Attendee"
    Vip TYPE = "Vip"
)

type Invite struct {
	gorm.Model
	UserID uint `json:"user_id" gorm:"not null"`
	User User `gorm:"foreignkey:UserId"`
	EventID uint `json:"event_id" gorm:"not null"`
	Event User `gorm:"foreignkey:EventId"`
	Type TYPE `gorm:"not null"`
}

func (i Invite) Validate() error {
	return validation.ValidateStruct(&i, 
		validation.Field(
			&i.UserID,
			validation.Required,
		),
		validation.Field(
			&i.EventID,
			validation.Required,
		),
		validation.Field(
			&i.Type,
			validation.Required,
			validation.In("Attendee", "Vip"),
		),
	
	)
}