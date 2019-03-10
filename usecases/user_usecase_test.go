package usecases

import (
	"github.com/habibridho/cleango/entities"
	"testing"
)

type MockDataSource struct{}

func (m *MockDataSource) GetUser(username string) (user entities.User, err error) {
	return entities.User{
		Username: "habibridho",
		Password: "123",
	}, nil
}

func TestUserUseCase_Login(t *testing.T) {
	mDataSource := &MockDataSource{}
	uc := NewUserUseCase(mDataSource)

	t.Run("Positive test", func(t *testing.T) {
		success, err := uc.Login("habibridho", "123")
		if err != nil {
			t.Fatalf("There shouldn't be any error. Got: %v", err)
		}

		if !success {
			t.Fatalf("Should be success")
		}
	})

	t.Run("Negative test", func(t *testing.T) {
		success, err := uc.Login("habibridho", "456")
		if err != nil {
			t.Fatalf("There shouldn't be any error. Got: %v", err)
		}

		if success {
			t.Fatalf("Should be fail")
		}
	})
}
