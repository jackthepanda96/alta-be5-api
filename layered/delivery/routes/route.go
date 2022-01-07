package routes

import (
	"Project/playground/be5/rest/layered/delivery/controllers/auth"
	"Project/playground/be5/rest/layered/delivery/controllers/user"
	custmiddlewares "Project/playground/be5/rest/layered/delivery/custMiddlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.GET("/users", uc.Get())

	e.POST("/login", ac.Login())

	e.POST("/users", uc.Insert(), custmiddlewares.JWTMiddleware())
}
