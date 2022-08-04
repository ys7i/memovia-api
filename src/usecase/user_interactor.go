package usecase

import "github.com/ys7i/memorizer/domain"

type UserInteractor struct {
	UserRepo UserRepository
}

func (interactor *UserInteractor) Get(id int) (user domain.User, resultStatus *ResultStatus) {
	user, err := interactor.UserRepo.FindById(id)
	if err!= nil {
		return domain.User{}, NewResultStatus(404, err)
	}
	return user, NewResultStatus(200, nil)
}