package address

import (
	"context"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID uint) ([]*Address, error)
}
