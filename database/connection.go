package database

import (
	"golang-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDatabase() {
	conn, err := gorm.Open(mysql.Open("root:@/golang_auth"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
