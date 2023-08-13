package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type TYPE string
type APPROVAL string

const (
	Attendee TYPE = "Attendee"
	Vip      TYPE = "Vip"
)

const (
	Pending APPROVAL = "Pending"
	Accept  APPROVAL = "Accept"
	Decline APPROVAL = "Decline"
)

type Invite struct {
	gorm.Model
	UserID   uint     `json:"user_id" gorm:"not null"`
	User     User     `gorm:"foreignkey:UserID"`
	EventID  uint     `json:"event_id" gorm:"not null"`
	Event    User     `gorm:"foreignkey:EventID"`
	Type     TYPE     `gorm:"not null"`
	Approval APPROVAL `json:"approval" gorm:"not null;default:'Pending'"`
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
		validation.Field(
			&i.Approval,
			validation.In("Pending", "Accept", "Decline"),
		),
	)
}
