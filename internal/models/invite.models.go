package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type TYPE string
type APPROVAL string

const (
	ATTENDEE TYPE = "ATTENDEE"
	VIP      TYPE = "VIP"
)

const (
	PENDING APPROVAL = "PENDING"
	ACCEPT  APPROVAL = "ACCEPT"
	DECLINE APPROVAL = "DECLINE"
)

type Invite struct {
	gorm.Model
	UserID   uint     `json:"user_id" gorm:"not null"`
	User     User     `gorm:"foreignkey:UserID"`
	EventID  uint     `json:"event_id" gorm:"not null"`
	Event    Event    `gorm:"foreignkey:EventID"`
	Type     TYPE     `gorm:"not null"`
	Approval APPROVAL `json:"approval" gorm:"not null;default:'PENDING'"`
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
			validation.In("ATTENDEE", "VIP"),
		),
		validation.Field(
			&i.Approval,
			validation.In("PENDING", "ACCEPT", "DECLINE"),
		),
	)
}
