package routes

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// localhost:8000/users/
	e.Pre(middleware.RemoveTrailingSlash()) //-->> proses dulu / paling akhir dihilangkan
	// lanjut ke e.GET
	// e.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, path=${path} status=${status}\n",
	// }))

	// e.Use(middleware.RemoveTrailingSlash())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, path=${path} status=${status} use\n",
	// }))

	// custMiddlewares.CustomMiddleware(e)
	eAuth := e.Group("/akses")
	eAuth.Use(middleware.JWT([]byte(configs.JWT_SECRET)))
	eAuth.POST("/users", controllers.CreateUser)
	// eAuth.Use(middleware.BasicAuth(custMiddlewares.DBBasicAuth))
	// eAuth.GET("/users", controllers.GetAllAccount)

	// e.Use(middleware.JWT(configs.JWT_SECRET))

	e.GET("/users/:id", controllers.CheckParameter)

	e.GET("/users", controllers.GetAllUser)
	e.POST("/login", controllers.Login)

	return e
}
