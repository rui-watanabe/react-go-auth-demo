package database

import (
	"react-go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root@tcp(localhost:3306)/react-go-auth"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connection mysql")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
