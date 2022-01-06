package main

import (
	"Project/playground/be5/rest/layered/configs"
	"Project/playground/be5/rest/layered/delivery/controllers/user"
	"Project/playground/be5/rest/layered/delivery/routes"
	_userRepo "Project/playground/be5/rest/layered/repository/user"
	"Project/playground/be5/rest/layered/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	userController := user.New(userRepo)

	e := echo.New()

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
