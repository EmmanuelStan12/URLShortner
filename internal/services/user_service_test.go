package services

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"testing"
)

func InitMockUsers(mock sqlmock.Sqlmock) {
	password, _ := util.HashPassword("test.password")

	users := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "test.name", "test@email.com", password)

	expectedSQL := "SELECT \\* FROM \\`users\\` WHERE \\(email = \\? AND password = \\?\\) AND \\`users\\`.\\`deleted_at\\` IS NULL"
	mock.ExpectQuery(expectedSQL).WillReturnRows(users)
}

func TestUserService_Login(t *testing.T) {
	sqlDb, db, mockDB := InitDBMock(t)
	defer sqlDb.Close()
	InitMockUsers(mockDB)
	userService := UserService{DB: db}

	t.Run("login existing user", func(t *testing.T) {
		request := dto.LoginUserRequest{Email: "test@email.com", Password: "test.password"}

		result := userService.Login(request)

		if result.Email != request.Email {
			t.Errorf("Expected %s, but got %s for logging in user", request.Email, result.Email)
		}
	})

	t.Run("Panic on empty password", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				_, ok := r.(*apperrors.Error)
				if !ok {
					t.Errorf("expected an error, got %v", r)
				}
			} else {
				t.Error("expected panic but did not panic")
			}
		}()

		request := dto.LoginUserRequest{Email: "test@email.com", Password: ""}
		userService.Login(request)
	})

	t.Run("Panic on empty email", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				_, ok := r.(*apperrors.Error)
				if !ok {
					t.Errorf("expected an error, got %v", r)
				}
			} else {
				t.Error("expected panic but did not panic")
			}
		}()

		request := dto.LoginUserRequest{Email: "", Password: "test.password"}
		userService.Login(request)
	})

	t.Run("Panic on non-existing user", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				_, ok := r.(error)
				if !ok {
					t.Errorf("expected an error, got %v", r)
				}
			} else {
				t.Error("expected panic but did not panic")
			}
		}()

		request := dto.LoginUserRequest{Email: "test1@email.com", Password: "test.password"}
		userService.Login(request)
	})
}
