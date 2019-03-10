package usecases

import (
	"errors"
	"github.com/habibridho/cleango/entities"
)

type UserData interface {
	GetUser(username string) (user entities.User, err error)
}

type userUseCase struct {
	Data UserData
}

func NewUserUseCase(ud UserData) *userUseCase {
	return &userUseCase{Data: ud}
}

func (uc *userUseCase) Login(username string, password string) (success bool, err error) {
	user, err := uc.Data.GetUser(username)
	if err != nil {
		return
	}

	if user.Password != password {
		err = errors.New("Password not match")
		return
	}

	success = true
	return
}
