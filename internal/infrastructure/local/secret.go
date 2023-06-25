package local

import (
	"context"
	"os"

	"github.com/taniko/nullable"
	"github.com/taniko/rin/internal/domain/repository"
)

func NewSecretClient(ctx context.Context) (repository.Secret, error) {
	return &secretManagerClient{}, nil
}

type secretManagerClient struct {
}

func (s secretManagerClient) GetSecret(_ context.Context, name string) (string, error) {
	var key nullable.Nullable[string]
	switch name {
	case "rin-db-password":
		key = nullable.New("DB_PASSWORD")
	case "signing-key":
		key = nullable.New("SIGNING_KEY")
	}
	if key.IsNull() {
		return "", repository.ErrNotFound
	}
	if v, ok := os.LookupEnv(key.Value()); ok {
		return v, nil
	}
	return "", repository.ErrNotFound
}
