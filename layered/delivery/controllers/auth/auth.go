package auth

import (
	"Project/playground/be5/rest/layered/repository/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}

		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, "Kesalahan input file")
		}

		checkedUser, err := a.repo.Login(loginFormat.Name, loginFormat.HP)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "there is something wrong in the system")
		}

		token, err := createToken(checkedUser.HP)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, "cannot process obtained value")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"data":    checkedUser,
			"token":   token,
		})
	}
}

func createToken(hp string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = 1
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["name"] = hp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RAHASIA"))
}
