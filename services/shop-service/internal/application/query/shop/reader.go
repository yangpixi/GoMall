package shop

import "context"

type Reader interface {
	GetShopID(ctx context.Context, userID int64) (*VO, error)
}
