package main

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"Project/playground/be5/rest/echo-structure/routes"
)

func main() {
	configs.InitDB()

	app := routes.New()

	app.Start(":8000")
}
