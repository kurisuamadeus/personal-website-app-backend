package helper

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matches := pattern.MatchString(email)
	if !matches {
		return errors.New("email is not valid")
	}

	return nil
}
