package validator

import (
	"regexp"
	"unicode/utf8"
)

func IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return len(email) <= 254 && emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	return utf8.RuneCountInString(password) < 5
}
