package account

import (
	"time"
)

type (
	ID          string
	Token       string
	Email       string
	Password    string
	RawPassword string
)

type Account struct {
	ID          string
	Name        string
	DisplayName string
	Email       Email
	Password    Password
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
