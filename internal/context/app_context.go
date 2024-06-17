package context

import (
	"github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/database"
	"github.com/EmmanuelStan12/URLShortner/internal/services"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	"github.com/EmmanuelStan12/URLShortner/pkg/jwt"
	"gorm.io/gorm"
)

// Context These are the app dependencies, majorly the singletons
type Context struct {
	Config     *config.Config
	DB         *gorm.DB
	JWTService jwt.JWTService
	Routes     util.Routes
}

func (c Context) GetUserService() services.IUserService {
	return services.UserService{DB: c.DB}
}

func InitContext() (*Context, error) {
	c, err := config.InitRootConfig()
	if err != nil {
		return nil, err
	}

	s := jwt.JWTService{
		SecretKey: c.Security.JWT.SecretKey,
		Issuer:    c.Security.JWT.Issuer,
	}

	db, err := database.InitDatabase(c.DB)
	if err != nil {
		return nil, err
	}
	context := &Context{
		Config:     c,
		DB:         db,
		JWTService: s,
		Routes:     util.InitRoutes(),
	}

	return context, nil
}
