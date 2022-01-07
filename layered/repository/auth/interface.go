package auth

import "Project/playground/be5/rest/layered/entities"

type Auth interface {
	Login(name, hp string) (entities.User, error)
}
