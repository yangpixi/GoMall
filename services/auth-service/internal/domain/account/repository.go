package account

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id uint) (*Account, error)
	FindByUsername(ctx context.Context, username string) (*Account, error)
}
