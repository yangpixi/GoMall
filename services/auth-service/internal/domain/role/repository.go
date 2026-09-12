package role

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID uint) ([]*Role, error)
	GrantByUserID(ctx context.Context, userID uint, roleID []uint) error
}
