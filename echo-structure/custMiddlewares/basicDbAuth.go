package custMiddlewares

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

func DBBasicAuth(username, password string, c echo.Context) (bool, error) {
	checkUser := models.Account{Username: username, Location: password}
	if err := configs.DB.Where("username = ? AND location = ?", checkUser.Username, checkUser.Location).Find(&checkUser).Error; err != nil || checkUser.ID == 0 {
		return false, nil
	}
	fmt.Println(checkUser)
	return true, nil
}
