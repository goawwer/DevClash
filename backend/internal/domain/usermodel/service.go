package usermodel

import (
	"errors"
	"net/mail"
	"regexp"
)

var (
	ErrRequiredFields  = errors.New("these fields required and cannot be empty")
	ErrInvalidEmail    = errors.New("email is invalid")
	ErrInvalidUsername = errors.New("username must contains only letters and number and be at least 5 characters")
)

func Validate(u *User) error {
	if u.Email == "" || u.Username == "" {
		return ErrRequiredFields
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	usernameRegex := regexp.MustCompile("[a-zA-Z0-9]+")
	if len(u.Username) < 5 || len(u.Username) > 64 || !usernameRegex.MatchString(u.Username) {
		return ErrInvalidUsername
	}

	return nil
}
