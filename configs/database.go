package configs

import (
	"go-auth-jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_auth_jwt?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connected")
}
