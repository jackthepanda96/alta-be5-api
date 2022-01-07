package user

import "Project/playground/be5/rest/layered/entities"

type CreateUserRequest struct {
	Nama string `json:"nama" form:"nama"`
	HP   string `json:"hp" form:"hp"`
}

type GetUserResponseFormat struct {
	Data    []entities.User `json:"data"`
	Message string          `json:"message"`
}
