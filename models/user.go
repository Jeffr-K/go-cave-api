package models

import "time"

type User struct {
	Id        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
