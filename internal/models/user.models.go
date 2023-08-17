package models

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// type GENDER string

// const (
// 	MALE   GENDER = "MALE"
// 	FEMALE GENDER = "FEMALE"
// )

type User struct {
	gorm.Model
	Name          string  `json:"name" gorm:"not null"`
	Birthday      string  `json:"birthday" gorm:"not null"`
	Email         string  `json:"email" gorm:"unique_index;not null"`
	Password      string  `json:"password" gorm:"not null"`
	Gender        string  `json:"gender" gorm:"not null"`
	CreatedEvents []Event `gorm:"foreignkey:UserID"`
	Events        []Event `gorm:"many2many:user_events;"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Name,
			validation.Required,
		),
		validation.Field(
			&u.Birthday,
			validation.Required,
		),
		validation.Field(
			&u.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&u.Password,
			validation.Required,
		),
		validation.Field(
			&u.Gender,
			validation.Required,
			validation.In("FEMALE", "MALE"),
		),
	)
}

func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return errors.New("password should not be empty")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(h)

	return nil

}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
