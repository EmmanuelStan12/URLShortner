package services

import (
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"gorm.io/gorm"
	"testing"
)

func initTestUser(db *gorm.DB) *models.User {
	password, _ := util.HashPassword("Test.password.1")
	user := models.User{
		Name:     "test.name.1",
		Email:    "test1@email.com",
		Password: password,
	}
	db.Save(&user)
	return &user
}

func TestUserService_Login(t *testing.T) {
	db, _ := initTestDB(t)
	initTestUser(db)
	defer teardownTestDB(db)
	userService := UserService{DB: db}

	t.Run("Panic on empty password", func(t *testing.T) {
		handlePanic[*apperrors.Error](t, func() {
			request := dto.LoginUserRequest{Email: "test@email.com", Password: ""}
			userService.Login(request)
		})
	})

	t.Run("Panic on empty email", func(t *testing.T) {
		handlePanic[*apperrors.Error](t, func() {
			request := dto.LoginUserRequest{Email: "", Password: "test.password"}
			userService.Login(request)
		})
	})

	t.Run("Panic on non-existing user", func(t *testing.T) {
		handlePanic[*apperrors.Error](t, func() {
			request := dto.LoginUserRequest{Email: "test100@email.com", Password: "test.password"}
			userService.Login(request)
		})
	})
}

func TestUserService_GetById(t *testing.T) {
	db, _ := initTestDB(t)
	defer teardownTestDB(db)
	userService := UserService{DB: db}

	t.Run("get existing user", func(t *testing.T) {
		user := models.User{
			Name:     fmt.Sprintf("test.name.%d", 1),
			Email:    fmt.Sprintf("test%d@email.com", 1),
			Password: "test.password",
		}
		db.Save(&user)

		result := userService.GetById(user.ID)

		if result.ID != 1 {
			t.Errorf("Expected %d, but got %d for getting an existing user", 1, result.ID)
		}
	})
}

func TestUserService_Create(t *testing.T) {
	db, _ := initTestDB(t)
	initTestUser(db)
	defer teardownTestDB(db)
	userService := UserService{DB: db}

	t.Run("test create user panic tests", func(t *testing.T) {
		panicTests := []struct {
			request  dto.RegisterUserRequest
			name     string
			expected any
		}{
			{
				request: dto.RegisterUserRequest{
					Email:    "test4@email.com",
					Password: "Test.password.123!",
				},
				name: "create user without name",
			},
			{
				request: dto.RegisterUserRequest{
					Name:     "test.name",
					Password: "Test.password.123!",
				},
				name: "create user without email",
			},
			{
				request: dto.RegisterUserRequest{
					Name:     "test.name",
					Email:    "test.email.com",
					Password: "Test.password.123!",
				},
				name: "create user with invalid email",
			},
			{
				request: dto.RegisterUserRequest{
					Name:     "test.name",
					Email:    "test1@email.com",
					Password: "Test.password.123!",
				},
				name: "create user with already existing email",
			},
			{
				request: dto.RegisterUserRequest{
					Name:  "test.name",
					Email: "test1@email.com",
				},
				name: "create user with empty password",
			},
			{
				request: dto.RegisterUserRequest{
					Name:     "test.name",
					Email:    "test1@email.com",
					Password: "test.password",
				},
				name: "create user with invalid password",
			},
		}

		for _, tt := range panicTests {
			t.Run(tt.name, func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					userService.Create(&tt.request)
				})
			})
		}
	})

	t.Run("create user", func(t *testing.T) {
		request := dto.RegisterUserRequest{
			Name:     "test.name",
			Email:    "test5@email.com",
			Password: "Test.Password!1",
		}

		userService.Create(&request)
	})
}

func TestUserService_Update(t *testing.T) {
	db, _ := initTestDB(t)
	user := initTestUser(db)
	defer teardownTestDB(db)
	userService := UserService{DB: db}

	t.Run("test create user panic tests", func(t *testing.T) {
		panicTests := []struct {
			request dto.UpdateUserRequest
			name    string
		}{
			{
				request: dto.UpdateUserRequest{
					RegisterUserRequest: dto.RegisterUserRequest{
						Email: "test1.email.com",
					},
				},
				name: "update user with invalid email",
			},
			{
				request: dto.UpdateUserRequest{
					RegisterUserRequest: dto.RegisterUserRequest{
						Email: "test1@email.com",
					},
				},
				name: "update user with already existing email",
			},
			{
				request: dto.UpdateUserRequest{
					RegisterUserRequest: dto.RegisterUserRequest{
						Password: "test.password",
					},
				},
				name: "update user with invalid new password",
			},
			{
				request: dto.UpdateUserRequest{
					RegisterUserRequest: dto.RegisterUserRequest{
						Password: "Test.password.1",
					},
					OldPassword: "Test.password.5",
				},
				name: "update user with invalid old password",
			},
		}

		for _, tt := range panicTests {
			t.Run(tt.name, func(t *testing.T) {
				handlePanic[*apperrors.Error](t, func() {
					userService.Update(user.ID, &tt.request)
				})
			})
		}
	})

	t.Run("update user", func(t *testing.T) {
		request := dto.UpdateUserRequest{
			RegisterUserRequest: dto.RegisterUserRequest{
				Name:     "test.name.updated",
				Email:    "test.updated@email.com",
				Password: "Test.password.updated.1",
			},
			OldPassword: "Test.password.1",
		}

		userService.Update(user.ID, &request)
	})
}

func TestUserService_Delete(t *testing.T) {
	db, _ := initTestDB(t)
	user := initTestUser(db)
	defer teardownTestDB(db)
	userService := UserService{DB: db}

	t.Run("delete non-existent user", func(t *testing.T) {
		handlePanic[*apperrors.Error](t, func() {
			userService.Delete(user.ID + 1)
		})
	})

	t.Run("delete user", func(t *testing.T) {
		userService.Delete(user.ID)
		handlePanic[*apperrors.Error](t, func() {
			userService.GetById(user.ID + 1)
		})
	})
}
