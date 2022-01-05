package database

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/models"
)

func CreateUser(u models.User) (interface{}, error) {
	// tmp := models.User{}
	if err := configs.DB.Save(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func GetAllUser() ([]models.User, error) {
	arrUser := []models.User{}

	if err := configs.DB.Find(&arrUser).Error; err != nil {
		return nil, err
	}
	return arrUser, nil
}

func Login(email, password string) (models.User, error) {
	foundUser := models.User{}
	if err := configs.DB.Where("email = ? AND password = ?", email, password).Find(&foundUser).Error; err != nil {
		return foundUser, err
	}
	return foundUser, nil
}
