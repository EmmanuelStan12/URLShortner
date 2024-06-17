package handlers

import (
	"encoding/json"
	"github.com/EmmanuelStan12/URLShortner/internal/constants"
	"github.com/EmmanuelStan12/URLShortner/internal/context"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(context.Context)
	userService := ctx.GetUserService()
	request := dto.LoginUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(apperrors.BadRequestError(err))
	}
	user := userService.Login(request)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	token, err := ctx.JWTService.GenerateToken(user.ID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(dto.SuccessResponse[any]{
		Status: http.StatusOK,
		Data: struct {
			User  dto.UserDTO `user:"json"`
			Token string      `token:"json"`
		}{User: user, Token: token},
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(constants.AppContextKey).(context.Context)
	userService := ctx.GetUserService()
	request := dto.RegisterUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(apperrors.BadRequestError(err))
	}
	user := userService.Create(&request)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	token, err := ctx.JWTService.GenerateToken(user.ID)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(dto.SuccessResponse[any]{
		Status: http.StatusOK,
		Data: struct {
			User  dto.UserDTO `user:"json"`
			Token string      `token:"json"`
		}{User: user, Token: token},
	})
}
