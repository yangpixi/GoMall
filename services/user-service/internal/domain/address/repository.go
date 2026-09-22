package address

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Address, error)
	FindByUserID(ctx context.Context, userID int64) ([]*Address, error)
	Save(ctx context.Context, address *Address) error
	Delete(ctx context.Context, id, userID int64) error
}
