package user

import (
	"Project/playground/be5/rest/layered/repository/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo user.UserRepository
}

func New(user user.UserRepository) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something wrong")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    users,
		})
	}
}
