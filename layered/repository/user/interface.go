package user

import "Project/playground/be5/rest/layered/entities"

type User interface {
	Get() ([]entities.User, error)
}
