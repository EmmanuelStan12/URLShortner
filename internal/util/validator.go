package util

import (
	"net/mail"
	"unicode"
)

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsPasswordValid(password string) bool {
	lower := false
	number := false
	upper := false
	special := false
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) && c != ' ':
			lower = true
		default:
			return false
		}
	}
	return number && special && upper && lower && len(password) > 7
}
