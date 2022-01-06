package user

import (
	"Project/playground/be5/rest/layered/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get() ([]entities.User, error) {
	users := []entities.User{}
	if err := ur.db.Find(&users).Error; err != nil {
		log.Warn("Found database error:", err)
		return nil, err
	}
	return users, nil
}
