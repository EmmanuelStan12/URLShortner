package util

const (
	R_LOGIN    = "/login"
	R_REGISTER = "/register"
)

type Routes struct {
	PublicRoutes []string
}

func InitRoutes() Routes {
	return Routes{
		PublicRoutes: []string{R_LOGIN, R_REGISTER},
	}
}

func (r Routes) IsPublic(route string) bool {
	for _, publicRoute := range r.PublicRoutes {
		if publicRoute == route {
			return true
		}
	}
	return false
}
