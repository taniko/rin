package model

import "time"

type Account struct {
	ID          string
	Name        string
	DisplayName string `db:"display_name"`
	Email       string
	Password    string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
