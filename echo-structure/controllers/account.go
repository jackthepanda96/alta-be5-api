package controllers

import (
	"Project/playground/be5/rest/echo-structure/lib/database"
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

// Bagian dimana kita menerima request
// Kalian dapat mengolah dulu data-data dalam parameter yang dikirim lewat request
// Jika data dirasa telah siap. kalian bisa mengambil data ke database lewat fungsi
// pada bagian database

func GetAllAccount(c echo.Context) error {
	res, err := database.GetAllAccount()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan pada server")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    res,
	})
}

func CheckParameter(c echo.Context) error {
	sid := c.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error convert string")
	}
	return c.String(http.StatusOK, "Hasil get parameter bisa diconvert ke int"+string(id))
}
