package entities

import (
	"time"
)

type RegisterRequest struct {
	Email          string    `json:"email" validate:"required,email,gte=10,lte=255"`
	DisplayName    string    `json:"display_name" validate:"required,gte=1,lte=32"`
	Username       string    `json:"username" validate:"required,ascii,gte=6,lte=255"`
	Password       string    `json:"password" validate:"required,ascii,gte=6,lte=255"`
	DateOfBirthDay time.Time `json:"date_of_birth_day" validate:"required"`
}
