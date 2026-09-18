package role

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID int64) ([]*Role, error)
	GrantByUserID(ctx context.Context, userID int64, roleID []int64) error
}
