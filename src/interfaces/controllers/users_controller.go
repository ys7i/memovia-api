package controllers

import (
	"strconv"

	"github.com/ys7i/memorizer/src/interfaces/database"
	"github.com/ys7i/memorizer/src/usecase"
)



type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db *database.DBRepository) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
				UserRepo: db.UserRepo,  
		},
	}
}

func (controller *UsersController) Get(c Context) {

    id, _ := strconv.Atoi(c.Param("id"))

    user, res := controller.Interactor.Get(id)
    if res.Error != nil {
        c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
        return
    }
    c.JSON(res.StatusCode, NewH("success", user))
}