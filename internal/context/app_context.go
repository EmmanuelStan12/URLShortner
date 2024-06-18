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

func (c Context) GetUrlService() services.IUrlService {
	return services.UrlService{DB: c.DB}
}

func InitRootContext() (*Context, error) {
	c, err := config.InitRootConfig()
	if err != nil {
		return nil, err
	}
	ctx, err := InitContext(c)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func InitContext(conf *config.Config) (*Context, error) {
	s := jwt.JWTService{
		SecretKey: conf.Security.JWT.SecretKey,
		Issuer:    conf.Security.JWT.Issuer,
	}

	db, err := database.InitDatabase(conf.DB)
	if err != nil {
		return nil, err
	}
	context := &Context{
		Config:     conf,
		DB:         db,
		JWTService: s,
		Routes:     util.InitRoutes(),
	}

	return context, nil
}
