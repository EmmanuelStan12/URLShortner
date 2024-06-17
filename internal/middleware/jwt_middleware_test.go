package middleware

import (
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	"github.com/EmmanuelStan12/URLShortner/pkg/jwt"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetById(id uint) dto.UserDTO {
	user := dto.UserDTO{
		Name:      "test.name",
		Email:     "test.email",
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return user
}

func (m *MockUserService) Update(id uint, request *dto.UpdateUserRequest) dto.UserDTO {
	user := dto.UserDTO{
		Name:      request.Name,
		Email:     request.Email,
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return user
}

func (m *MockUserService) Create(request *dto.RegisterUserRequest) dto.UserDTO {
	user := dto.UserDTO{
		Name:      request.Name,
		Email:     request.Email,
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return user
}

func (m *MockUserService) Delete(id uint) dto.UserDTO {
	user := dto.UserDTO{
		Name:      "test.name",
		Email:     "test.email",
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return user
}

func (m *MockUserService) Login(request dto.LoginUserRequest) dto.UserDTO {
	user := dto.UserDTO{
		Name:      "test.name",
		Email:     "test.email",
		ID:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	return user
}

func TestJWTMiddleware(t *testing.T) {
	mockUserService := MockUserService{}
	jwtService := jwt.JWTService{
		SecretKey: "secret_key",
		Issuer:    "issuer",
	}
	errorMiddleware := ErrorMiddleware()
	jwtMiddleware := JWTMiddleware(jwtService, util.InitRoutes(), &mockUserService)

	token, err := jwtService.GenerateToken(1)
	if err != nil {
		t.Errorf("Error occurred while generating token, %v.", err)
		return
	}

	tests := []struct {
		name           string
		token          string
		expectedStatus int
	}{
		{
			name:           "Valid token",
			token:          token,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid token",
			token:          "Invalid token",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Empty token",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "https://example.com", nil)
			w := httptest.NewRecorder()
			r.Header.Set("Authorization", "Bearer "+tt.token)

			handler := errorMiddleware(jwtMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})))

			handler.ServeHTTP(w, r)

			if tt.expectedStatus != w.Code {
				t.Errorf("Invalid status code, expected %d, got %d.", tt.expectedStatus, w.Code)
			}

			t.Logf("Response %v", w.Body)
		})
	}
}
