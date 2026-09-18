package address

import (
	"context"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID int64) ([]*Address, error)
}
