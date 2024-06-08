package app

import (
	"github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/database"
	"gorm.io/gorm"
)

// Context These are the app dependencies, majorly the singletons
type Context struct {
	Config *config.Config
	DB     *gorm.DB
}

func InitContext() (*Context, error) {
	c, err := config.InitRootConfig()
	if err != nil {
		return nil, err
	}

	db, err := database.InitDatabase(c.DB)
	context := &Context{
		Config: c,
		DB:     db,
	}

	return context, nil
}
