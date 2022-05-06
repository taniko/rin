package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/taniko/rin/internal/domain/model"
	"github.com/taniko/rin/internal/domain/repository"
	"github.com/taniko/sumire"
)

type account struct {
	db     *sqlx.DB
	logger *sumire.Sumire
}

func (a account) FindByNameOrEmail(ctx context.Context, name, email string) (*model.Account, error) {
	rows, err := a.db.NamedQuery(`SELECT * FROM account WHERE name = :name OR email = :email`, map[string]interface{}{
		"name":  name,
		"email": email,
	})
	if err != nil {
		return nil, err
	}
	var m model.Account
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			a.logger.ErrorContext(ctx, fmt.Sprintf("failed to close rows: %v", err), nil)
		}
	}(rows)
	if !rows.Next() {
		return nil, nil
	}
	if err := rows.StructScan(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a account) Create(_ context.Context, account model.Account) (*model.Account, error) {
	_, err := a.db.NamedExec(`INSERT INTO account (id, name, display_name, email, password) 
		VALUES (:id, :name, :display_name, :email, :password)`, account)
	if err != nil {
		return nil, err
	}
	return &account, err
}

func NewAccountDatabase(db *sqlx.DB, logger *sumire.Sumire) repository.Account {
	return account{
		db:     db,
		logger: logger,
	}
}
