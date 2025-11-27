package repository

import (
	"errors"

	"github.com/goawwer/devclash/internal/domain"
	"github.com/lib/pq"
)

var constraintMap = map[string]error{
	"accounts_email_key":  domain.ErrEmailTaken,
	"users_username_key":  domain.ErrUsernameTaken,
	"organizers_name_key": domain.ErrOrganizerNameTaken,
}

func mapError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			if domainErr, ok := constraintMap[pqErr.Constraint]; ok {
				return domainErr
			}
		}
	}
	return err
}
