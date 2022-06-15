package models

import "time"

type Address struct {
	Id        uint `gorm:"primaryKey"`
	Street    string
	Detail    string
	Zipcode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
