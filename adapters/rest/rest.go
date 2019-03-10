package rest

import (
	"github.com/habibridho/cleango/usecases"
	"net/http"
)

type RestAdapter struct {
	UserUC usecases.UserUseCase
}

func (r RestAdapter) Login(username string, password string) Response {
	if username == "" {
		return Response{
			Status:  http.StatusBadRequest,
			Message: "username is empty",
		}
	}

	if password == "" {
		return Response{
			Status:  http.StatusBadRequest,
			Message: "password is empty",
		}
	}

	success, err := r.UserUC.Login(username, password)
	if err != nil {
		return Response{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}

	if !success {
		return Response{
			Status:  http.StatusUnauthorized,
			Message: "username/password is wrong",
		}
	}

	return Response{
		Status:  http.StatusOK,
		Message: "success",
	}
}
