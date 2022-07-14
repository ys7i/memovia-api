package usecase

import (
	"github.com/ys7i/memorizer/src/domain"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	FindById(id int) (domain.User, error)
}
