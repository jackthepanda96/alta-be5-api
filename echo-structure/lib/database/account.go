package database

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/models"
	"fmt"
	// "Project/playground/be5/rest/echo-structure/models"
)

// Fungsi untuk berinteraksi dengan database
// Penamaan file dan nama tabel pada database dapat di samakan

func GetAllAccount() ([]models.Account, error) {
	var users []models.Account

	if err := configs.DB.Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}
