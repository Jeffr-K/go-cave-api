package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"microuser/models"
)

var Gorm *gorm.DB

func init() {
	dsn := "root:root@tcp(localhost:3306)/cave_user?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	database.AutoMigrate(&models.User{}, &models.Address{})

	Gorm = database
}
