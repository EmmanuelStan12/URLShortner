package util

const (
	PRODUCTION  = "prod"
	TEST        = "test"
	DEV         = "dev"
	ENVIRONMENT = "env"
)

type Environment struct {
	Env string
}

func (env Environment) IsProd() bool {
	return env.Env == PRODUCTION
}

func (env Environment) isDev() bool {
	return env.Env == DEV
}

func (env Environment) isTest() bool {
	return env.Env == TEST
}
