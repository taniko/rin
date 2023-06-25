//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
package repository

import "context"

type Secret interface {
	GetSecret(ctx context.Context, key string) (string, error)
}
