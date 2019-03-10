package usecases

import (
	"github.com/habibridho/cleango/entities"
)

type UserData interface {
	GetUser(username string) (user entities.User, err error)
}

type userUseCase struct {
	Data UserData
}

func NewUserUseCase(ud UserData) userUseCase {
	return userUseCase{Data: ud}
}

func (uc userUseCase) Login(username string, password string) (success bool, err error) {
	user, err := uc.Data.GetUser(username)
	if err != nil {
		return
	}

	if user.Password != password {
		success = false
		return
	}

	success = true
	return
}
