package database

import (
	"github.com/ys7i/memorizer/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) FindById(id int) (domain.User, error) {
	user := domain.User{}
	err := repo.DB.First(&user, id).Error
	return user, err
}

func (repo *UserRepository) Create(user domain.User) (domain.User, error) {
	err := repo.DB.Create(&user).Error
	return user, err
}