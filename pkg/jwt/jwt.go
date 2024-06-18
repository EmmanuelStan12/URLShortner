package jwt

import (
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func getUserId(c jwt.MapClaims) (uint, *errors.Error) {
	v, ok := c["userId"]
	if ok {
		switch v := v.(type) {
		case float64:
			userId := uint(v)
			return userId, nil
		case int:
			return uint(v), nil
		case uint:
			return v, nil
		default:
			return 0, errors.UnauthorizedError(fmt.Errorf("%d is invalid for key userId", v))
		}
	}

	return 0, errors.UnauthorizedError(fmt.Errorf("%d is invalid for key userId", v))
}

type IJWTService interface {
	GenerateToken(userId int) (token string, err *errors.Error)
	ParseToken(token string) (claims jwt.Claims, err *errors.Error)
}

type JWTService struct {
	SecretKey string
	Issuer    string
}

func (s *JWTService) GenerateToken(userId uint) (string, *errors.Error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
			"iss":    s.Issuer,
		},
	)

	tokenStr, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", errors.InternalServerError(err)
	}
	return tokenStr, nil
}

func (s *JWTService) ParseToken(tokenStr string) (uint, *errors.Error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return 0, errors.UnauthorizedError(err)
	}

	if !token.Valid {
		return 0, errors.UnauthorizedError(fmt.Errorf("invalid Token %s", tokenStr))
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.InternalServerError("can't pass claims")
	}
	userId, e := getUserId(claims)
	if err != nil {
		return 0, e
	}

	return userId, nil
}
