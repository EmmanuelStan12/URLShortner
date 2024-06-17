package services

import (
	"errors"
	"github.com/EmmanuelStan12/URLShortner/internal/dto"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	apperrors "github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"gorm.io/gorm"
)

type IUserService interface {
	GetById(uint) dto.UserDTO
	Update(uint, *dto.UpdateUserRequest) dto.UserDTO
	Create(request *dto.RegisterUserRequest) dto.UserDTO
	Delete(uint) dto.UserDTO
	Login(request dto.LoginUserRequest) dto.UserDTO
}

type UserService struct {
	DB *gorm.DB
}

func (u UserService) Login(request dto.LoginUserRequest) dto.UserDTO {
	var user models.User
	errs := make([]string, 0)
	if request.Password == "" {
		errs = append(errs, "password cannot be empty")
	}
	if request.Email == "" {
		errs = append(errs, "email cannot be empty")
	}
	if len(errs) > 0 {
		panic(apperrors.ValidationError(errs))
	}
	password, _ := util.HashPassword(request.Password)
	result := u.DB.Where("email = ? AND password = ?", request.Email, password).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return util.ToUserDTO(user)
}

func (u UserService) GetById(userId uint) dto.UserDTO {
	user := u.getById(userId)
	return util.ToUserDTO(*user)
}

func (u UserService) Update(userId uint, request *dto.UpdateUserRequest) dto.UserDTO {
	user := u.getById(userId)
	errs := make([]string, 0)
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		if !util.IsEmailValid(request.Email) {
			errs = append(errs, "invalid email")
		} else if u.checkIfEmailExists(request.Email) {
			errs = append(errs, "email already exists")
		} else {
			user.Email = request.Email
		}
	}
	if request.Password != "" {
		if !util.IsPasswordValid(request.Password) {
			errs = append(errs, "password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
		} else if !util.ComparePasswordHash(request.OldPassword, user.Password) {
			errs = append(errs, "invalid old password")
		}
	}
	if len(errs) > 0 {
		panic(apperrors.ValidationError(errs))
	}
	if request.Password != "" {
		password, err := util.HashPassword(request.Password)
		if err != nil {
			panic(err)
		}
		user.Password = password
	}
	result := u.DB.Save(user)
	if result.Error != nil {
		panic(apperrors.InternalServerError(result.Error))
	}

	return util.ToUserDTO(*user)
}

func (u UserService) Create(request *dto.RegisterUserRequest) dto.UserDTO {
	user := models.User{}
	errs := make([]string, 0)
	if request.Name == "" {
		errs = append(errs, "name cannot be empty")
	}
	if request.Email != "" {
		if !util.IsEmailValid(request.Email) {
			errs = append(errs, "invalid email")
		} else if u.checkIfEmailExists(request.Email) {
			errs = append(errs, "email already exists")
		}
	} else {
		errs = append(errs, "email cannot be empty")
	}
	if request.Password == "" {
		errs = append(errs, "password cannot be empty")
	} else if !util.IsPasswordValid(request.Password) {
		errs = append(errs, "password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}
	if len(errs) > 0 {
		panic(apperrors.ValidationError(errs))
	}

	password, err := util.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}
	user.Password = password
	user.Email = request.Email
	user.Name = request.Name

	result := u.DB.Save(&user)
	if result.Error != nil {
		panic(apperrors.InternalServerError(result.Error))
	}

	return util.ToUserDTO(user)
}

func (u UserService) Delete(userId uint) dto.UserDTO {
	user := u.getById(userId)
	u.DB.Unscoped().Delete(models.User{}, userId)
	return util.ToUserDTO(*user)
}

func (u UserService) getById(userId uint) *models.User {
	user := models.User{}
	res := u.DB.First(&user, userId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			panic(errors.New("user not found"))
		}
		panic(res.Error)
	}

	return &user
}

func (u UserService) checkIfEmailExists(email string) bool {
	var user models.User
	result := u.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return false
	}
	return true
}
