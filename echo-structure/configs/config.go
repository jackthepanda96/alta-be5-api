package configs

import (
	"Project/playground/be5/rest/echo-structure/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	JWT_SECRET = "r4H4S!@aa"
)

func InitDB() {
	connectionString := "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.User{})
}
