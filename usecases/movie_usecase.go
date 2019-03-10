package usecases

import (
	"errors"
	"github.com/habibridho/cleango/entities"
)

type MovieData interface {
	StartTransaction() error
	EndTransaction(commit bool) error
	GetUser(userID uint64) (user entities.User, err error)
	GetMovie(movieID uint64) (entities.Movie, error)
	GetMovies() ([]entities.Movie, error)
	SaveLike(userID uint64, movieID uint64, status bool) error
}

type movieUseCase struct {
	Data MovieData
}

func NewMovieUseCase(md MovieData) *movieUseCase {
	return &movieUseCase{Data: md}
}

func (uc *movieUseCase) GetMovieList() (data []entities.Movie, err error) {
	data, err = uc.Data.GetMovies()
	return
}

func (uc *movieUseCase) LikeDislike(userID uint64, movieID uint64, status bool) (err error) {
	if err = uc.Data.StartTransaction(); err != nil {
		return
	}
	defer func() {
		_ = uc.Data.EndTransaction(err == nil)
	}()

	user, err := uc.Data.GetUser(userID)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("invalid user")
		return
	}

	movie, err := uc.Data.GetMovie(movieID)
	if err != nil {
		return
	}
	if movie.ID == 0 {
		err = errors.New("invalid movie")
		return
	}

	if err = uc.Data.SaveLike(userID, movieID, status); err != nil {
		return
	}

	return
}
