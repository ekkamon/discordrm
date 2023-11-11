package models

import (
	"discordrm/api/interval/entities"
	"time"
)

// gorm model
type User struct {
	UID            int       `gorm:"primaryKey;unique;autoIncrement;index;notNull" json:"uid"`
	Email          string    `gorm:"unique;notNull" json:"email"`
	DisplayName    string    `gorm:"notNull" json:"display_name"`
	Username       string    `gorm:"unique;notNull" json:"username"`
	Password       string    `gorm:"notNull" json:"password"`
	DateOfBirthDay time.Time `gorm:"notNull" json:"date_of_birth_day"`
	CreatedAt      time.Time `gorm:"notNull" json:"created_at"`
	UpdatedAt      time.Time `gorm:"notNull" json:"updated_at"`
}

func (u *User) ToPublic() entities.UserPublicData {
	return entities.UserPublicData{
		UID:            u.UID,
		Email:          u.Email,
		DisplayName:    u.DisplayName,
		Username:       u.Username,
		DateOfBirthDay: u.DateOfBirthDay,
		CreatedAt:      u.CreatedAt,
	}
}
