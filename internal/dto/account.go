package dto

import (
	"github.com/taniko/rin/internal/domain/model/account"
)

type CreateAccount struct {
	Name                 string
	DisplayName          string
	Email                account.Email
	Password             account.RawPassword
	PasswordConfirmation account.RawPassword
}

type LoginRequest struct {
	Email    account.Email
	Password account.RawPassword
}
