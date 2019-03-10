package rest

import (
	"github.com/habibridho/cleango/usecases"
	"net/http"
)

type RestAdapter struct {
	UserUC  usecases.UserUseCase
	MovieUC usecases.MovieUseCase
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

func (r RestAdapter) GetMovies() Response {
	data, err := r.MovieUC.GetMovieList()
	if err != nil {
		return Response{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}

	movies := []Movie{}
	for _, d := range data {
		movies = append(movies, Movie{
			ID:       d.ID,
			Title:    d.Title,
			Director: d.Director,
		})
	}

	return Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    movies,
	}
}

func (r RestAdapter) ToggleLike(userID uint64, movieID uint64, status bool) Response {
	if err := r.MovieUC.LikeDislike(userID, movieID, status); err != nil {
		if err.Error() == usecases.ERR_INVALID_USER {
			return Response{
				Status:  http.StatusBadRequest,
				Message: "invalid user id",
			}
		}

		if err.Error() == usecases.ERR_INVALID_MOVIE {
			return Response{
				Status:  http.StatusBadRequest,
				Message: "invalid movie id",
			}
		}

		return Response{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}

	return Response{
		Status:  http.StatusOK,
		Message: "success",
	}
}
