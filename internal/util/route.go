package util

import (
	"net"
	"net/http"
)

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

func GetIPAddress(r http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}
