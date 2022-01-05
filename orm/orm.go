package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Account struct {
	gorm.Model
	// ID          int
	Username    string `json:"username" form:"username"`
	Displayname string `json:"displayname" form:"displayname"`
	Location    string `json:"location" form:"location"`
}

type User struct {
	// gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitDB() {
	connectionString := "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
}

func InitMigrate() {
	DB.AutoMigrate(&Account{})
	// DB.AutoMigrate(&User{})
}

func init() {
	fmt.Println("Init Dari init")
	InitDB()
	InitMigrate()
}

func GetUserController(c echo.Context) error {
	var users []Account

	if err := DB.Find(&users).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

func InsetUserController(c echo.Context) error {
	userBaru := Account{Username: "yoga555", Displayname: "Yogawah", Location: "Indonesia"}

	if err := DB.Create(&userBaru).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Format input kurang tepat")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "berhasil create user",
		"data":    userBaru,
	})
}

func GetTweet(c echo.Context) error {
	type Tweet struct {
		Tweet    string
		Createby int
	}

	arrTemp := []Tweet{}
	// Akses manual per row
	// db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	// tmpRow, err := DB.Table("tweet").Select("tweet, createby").Rows()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan pada server")
	// }
	// for tmpRow.Next() {
	// 	tmp := Tweet{}
	// 	if err := tmpRow.Scan(&tmp.Tweet, &tmp.Createby); err != nil {
	// 		fmt.Println(err)
	// 		return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan pada server")
	// 	}
	// 	arrTemp = append(arrTemp, tmp)
	// }
	// Akses dengan .Raw
	DB.Raw("SELECT tweet, createby FROM tweet").Scan(&arrTemp)
	// if err := DB.Find(&tmp).Select("tweet, createby").Error; err != nil {
	// 	fmt.Println(err)
	// 	return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan pada server")
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get tweet",
		"data":    arrTemp,
	})
}

func main() {
	e := echo.New()
	e.GET("/accounts", GetUserController)
	e.GET("/tweet", GetTweet)
	e.GET("/user", InsetUserController)

	e.Start(":8000")
}
