package controllers

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/custMiddlewares"
	"Project/playground/be5/rest/echo-structure/lib/database"
	"Project/playground/be5/rest/echo-structure/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	tmp := models.User{}

	c.Bind(&tmp)
	res, err := database.CreateUser(tmp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Input ada yang tidak sesuai")
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "data berhasil ditambahkan",
		"data":    res,
	})
}

func GetAllUser(c echo.Context) error {
	res, err := database.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Terjadi Kesalahan pada Sistem")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil get data",
		"data":    res,
	})
}

func Login(c echo.Context) error {
	tmpLogin := models.User{}
	c.Bind(&tmpLogin)
	res, err := database.Login(tmpLogin.Email, tmpLogin.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Data tidak ditemukan")
	}
	fmt.Println(res)
	res.Token, _ = custMiddlewares.CreateToken(int(res.ID), configs.JWT_SECRET)
	// if res.Token == "" {
	// 	res.Token, _ = custMiddlewares.CreateToken(int(res.ID), configs.JWT_SECRET)
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil login",
		"data":    res,
	})
}
