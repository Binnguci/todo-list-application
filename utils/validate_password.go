package utils

import (
	"errors"
	"regexp"
	"todo-app/constants"
)

func ValidatePassword(password string) error {
	if len(password) < 6 || len(password) > 20 {
		return errors.New(constants.PASSWORD_INVALID)
	}

	lowercaseRe := regexp.MustCompile(`[a-z]`)
	if !lowercaseRe.MatchString(password) {
		return errors.New(constants.PASSWORD_INVALID)
	}

	uppercaseRe := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRe.MatchString(password) {
		return errors.New(constants.PASSWORD_INVALID)
	}

	digitRe := regexp.MustCompile(`\d`)
	if !digitRe.MatchString(password) {
		return errors.New(constants.PASSWORD_INVALID)
	}

	specialCharRe := regexp.MustCompile(`[@$!%*?&]`)
	if !specialCharRe.MatchString(password) {
		return errors.New(constants.PASSWORD_INVALID)
	}

	return nil
}
