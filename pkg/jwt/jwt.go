package jwt

import (
	builtInErrors "errors"
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	jwt.MapClaims
}

func (c CustomClaims) GetUserId() (int, *errors.Error) {
	v, ok := c.MapClaims["userId"]
	if ok {
		userId, ok := v.(int)
		if ok {
			return userId, nil
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
		CustomClaims{
			jwt.MapClaims{
				"userId": userId,
				"exp":    time.Now().Add(time.Hour * 24).Unix(),
				"iss":    s.Issuer,
			},
		},
	)

	tokenStr, err := token.SignedString(s.SecretKey)
	if err != nil {
		return "", errors.InternalServerError(err)
	}
	return tokenStr, nil
}

func (s *JWTService) ParseToken(tokenStr string) (int, *errors.Error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return s.SecretKey, nil
	})

	if err != nil {
		return 0, errors.UnauthorizedError(err)
	}

	if !token.Valid {
		return 0, errors.UnauthorizedError(fmt.Errorf("invalid Token %s", tokenStr))
	}

	claims, ok := token.Claims.(CustomClaims)

	if !ok {
		return 0, errors.InternalServerError(builtInErrors.New("can't pass claims"))
	}
	userId, e := claims.GetUserId()
	if err != nil {
		return 0, e
	}

	return userId, nil
}
