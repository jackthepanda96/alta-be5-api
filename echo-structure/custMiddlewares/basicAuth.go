package custMiddlewares

import "github.com/labstack/echo/v4"

func BasicAuth(user, password string, e echo.Context) (bool, error) {
	if user == "admin" && password == "admin" {
		return true, nil
	}
	return false, nil
}
