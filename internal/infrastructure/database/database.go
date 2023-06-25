package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/taniko/rin/internal/app/env"
)

type Config struct {
	Host     string
	Database string
	User     string
	Password string
}

type Password string

func New(environment env.Environment, password Password) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
		environment.DB.User, password, environment.DB.Host, environment.DB.NAME),
	)
}
