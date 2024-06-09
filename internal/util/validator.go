package util

import (
	"net/mail"
	"regexp"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsPasswordValid(password string) bool {
	regex := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[!@#~$%^&*(),.?":{}|<>])[A-Za-z\d!@#~$%^&*(),.?":{}|<>]{8,}$`

	re := regexp.MustCompile(regex)
	if !re.MatchString(password) {
		return false
	}
	return true
}
