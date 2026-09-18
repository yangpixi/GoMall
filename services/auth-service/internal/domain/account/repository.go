package account

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Account, error)
	FindByUsername(ctx context.Context, username string) (*Account, error)
	Save(ctx context.Context, account *Account) error
}
