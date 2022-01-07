package main

import (
	"Project/playground/be5/rest/layered/configs"
	"Project/playground/be5/rest/layered/delivery/controllers/auth"
	"Project/playground/be5/rest/layered/delivery/controllers/user"
	"Project/playground/be5/rest/layered/delivery/routes"
	_authRepo "Project/playground/be5/rest/layered/repository/auth"
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
	authRepo := _authRepo.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.RegisterPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
