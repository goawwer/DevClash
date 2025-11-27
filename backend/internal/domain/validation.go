package domain

import (
	"net/mail"
	"regexp"

	"github.com/goawwer/devclash/internal/dto"
)

func Validate(a dto.SignUpForm) error {
	if a.Username != "" {
		usernameRegex := regexp.MustCompile("[a-zA-Z0-9]+")

		if len(a.Username) < 5 || len(a.Username) > 64 || !usernameRegex.MatchString(a.Username) {
			return ErrInvalidUsername
		}
	}

	if a.Name != "" {
		organizerRegex := regexp.MustCompile("[a-zA-Z]+")

		if !organizerRegex.MatchString(a.Name) {
			return ErrInvalidCompanyName
		}
	}

	if a.Email == "" || a.Password == "" {
		return ErrRequiredFields
	}

	_, err := mail.ParseAddress(a.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
