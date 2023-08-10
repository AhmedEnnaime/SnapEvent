package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
) 

type UserEvent struct {
	gorm.Model
	UserID uint`json:"user_id"`
	EventID uint `json:"event_id"`
}

func (ue UserEvent) Validate() error {
	return validation.ValidateStruct(&ue,
		validation.Field(
			&ue.UserID,
			validation.Required,
		),
		validation.Field(
			&ue.EventID,
			validation.Required,
		),
	
	)
}