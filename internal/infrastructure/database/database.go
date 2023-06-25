package database

import (
	"time"

	"github.com/go-sql-driver/mysql"
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
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	conf := mysql.Config{
		User:      environment.DB.User,
		Passwd:    string(password),
		Addr:      environment.DB.Host,
		DBName:    environment.DB.NAME,
		ParseTime: true,
		Loc:       loc,
	}
	return sqlx.Connect("mysql", conf.FormatDSN())
}
