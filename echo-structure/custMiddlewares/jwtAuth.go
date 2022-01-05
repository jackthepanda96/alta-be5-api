package custMiddlewares

import (
	"Project/playground/be5/rest/echo-structure/configs"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id int, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractToken(e echo.Context) int {
	token := e.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]
		switch id.(type) {
		case float64:
			return int(id.(float64))
		default:
			return id.(int)
		}
	}
	return -1 //invalid
}
