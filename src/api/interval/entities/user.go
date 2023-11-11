package entities

import (
	"time"
)

type UserRegisterRequest struct {
	Email          string    `json:"email" validate:"required,email,gte=10,lte=255"`
	DisplayName    string    `json:"display_name" validate:"required,ascii,gte=1,lte=32"`
	Username       string    `json:"username" validate:"required,ascii,gte=6,lte=255"`
	Password       string    `json:"password" validate:"required,ascii,gte=6,lte=255"`
	DateOfBirthDay time.Time `json:"date_of_birth_day" validate:"required"`
}

type UserPublicData struct {
	UID            int       `json:"uid"`
	Email          string    `json:"email"`
	DisplayName    string    `json:"display_name"`
	Username       string    `json:"username"`
	DateOfBirthDay time.Time `json:"date_of_birth_day"`
	CreatedAt      time.Time `json:"created_at"`
}
