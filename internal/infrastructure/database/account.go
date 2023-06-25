package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	account2 "github.com/taniko/rin/internal/domain/model/account"
	"github.com/taniko/rin/internal/domain/repository"
	"golang.org/x/xerrors"
)

type account struct {
	db *sqlx.DB
}

func (a account) FindByID(ctx context.Context, ID account2.ID) (*account2.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccount(db *sqlx.DB) repository.Account {
	return account{
		db: db,
	}
}

func (a account) FindByNameOrEmail(ctx context.Context, name string, email account2.Email) (*account2.Account, error) {
	v, err := a.find(ctx, `SELECT * FROM account WHERE name = :name OR email = :email`, map[string]any{
		"name":  name,
		"email": email,
	})
	if err != nil {
		return nil, xerrors.Errorf("find account by name or email: %w", err)
	}
	return v, nil
}

func (a account) FindByEmail(ctx context.Context, email account2.Email) (*account2.Account, error) {
	v, err := a.find(ctx, `SELECT * FROM account WHERE email = :email`, map[string]any{
		"email": string(email),
	})
	if err != nil {
		return nil, xerrors.Errorf("find account by name: %w", err)
	}
	return v, nil
}

func (a account) find(ctx context.Context, query string, params map[string]any) (*account2.Account, error) {
	rows, err := a.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		return nil, xerrors.Errorf("find query: %w", err)
	}
	var m account2.Account
	defer func(rows *sqlx.Rows) {
		if err := rows.Close(); err != nil {
			log.Err(err).Msgf("failed to close rows")
		}
	}(rows)
	if !rows.Next() {
		return nil, repository.ErrNotFound
	}
	if err := rows.StructScan(&m); err != nil {
		return nil, xerrors.Errorf("row to account struct: %w", err)
	}
	return &m, nil
}

func (a account) Create(ctx context.Context, account account2.Account) (*account2.Account, error) {
	_, err := a.db.NamedExecContext(ctx, `INSERT INTO account (id, name, display_name, email, password) 
		VALUES (:id, :name, :display_name, :email, :password)`, account)
	if err != nil {
		return nil, xerrors.Errorf("create account: %w", err)
	}
	return &account, nil
}
