package query

import "context"

type Reader interface {
	GetOrder(ctx context.Context, id int64) (*VO, error)
}
