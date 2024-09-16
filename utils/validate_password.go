package utils

import (
	"errors"
	"regexp"
	"todo-app/constants"
)

func ValidatePassword(password string) error {
	const passwordPattern = constants.REGEX_PASSWORD
	re := regexp.MustCompile(passwordPattern)
	if !re.MatchString(password) {
		return errors.New(constants.PASSWORD_INVALID)
	}
	return nil
}
