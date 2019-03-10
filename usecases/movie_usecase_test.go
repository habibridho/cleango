package usecases

import (
	"github.com/habibridho/cleango/entities"
	"testing"
)

type MockMovieDataSource struct{}

func (m *MockMovieDataSource) StartTransaction() error {
	return nil
}
func (m *MockMovieDataSource) EndTransaction(commit bool) error {
	return nil
}
func (m *MockMovieDataSource) GetUser(userID uint64) (user entities.User, err error) {
	user = entities.User{ID: 1}
	return
}
func (m *MockMovieDataSource) GetMovie(movieID uint64) (entities.Movie, error) {
	return entities.Movie{ID: 1}, nil
}
func (m *MockMovieDataSource) GetMovies() ([]entities.Movie, error) {
	return []entities.Movie{}, nil
}
func (m *MockMovieDataSource) SaveLike(userID uint64, movieID uint64, status bool) error {
	return nil
}

func TestMovieUseCase_GetMovieList(t *testing.T) {
	mDataSource := &MockMovieDataSource{}
	uc := NewMovieUseCase(mDataSource)

	_, err := uc.GetMovieList()
	if err != nil {
		t.Fatalf("There should be no error, got: %v", err)
	}
}

func TestMovieUseCase_LikeDislike(t *testing.T) {
	mDataSource := &MockMovieDataSource{}
	uc := NewMovieUseCase(mDataSource)

	err := uc.LikeDislike(1, 1, true)
	if err != nil {
		t.Fatalf("There should be no error, got: %v", err)
	}
}
