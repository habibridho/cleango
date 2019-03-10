package usecases

import (
	"github.com/habibridho/cleango/entities"
	"github.com/kataras/go-errors"
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

func (uc *userUseCase) login(username string, password string) (success bool, err error) {
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
