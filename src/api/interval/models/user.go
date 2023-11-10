package models

import (
	"time"
)

// gorm model
type User struct {
	UID            int       `gorm:"primaryKey;unique;autoIncrement;index;notNull"`
	Email          string    `gorm:"unique;notNull"`
	DisplayName    string    `gorm:"notNull"`
	Username       string    `gorm:"unique;notNull"`
	Password       string    `gorm:"notNull"`
	DateOfBirthDay time.Time `gorm:"notNull"`
	CreatedAt      time.Time `gorm:"notNull"`
	UpdatedAt      time.Time `gorm:"notNull"`
}
