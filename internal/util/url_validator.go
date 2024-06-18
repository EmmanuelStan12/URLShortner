package util

import (
	"net/http"
	"net/url"
	"strings"
)

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func IsValidUrl(str string) bool {
	if !isUrl(str) {
		return false
	}
	resp, err := http.Get(str)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	// Check if the HTTP response is a success (2xx) or success-like code (3xx)
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return true
	}

	return false
}

func NormalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	u.Host = strings.ToLower(strings.TrimPrefix(u.Host, "www."))

	u.Path = strings.TrimSuffix(u.Path, "/")

	return u.String(), nil
}
