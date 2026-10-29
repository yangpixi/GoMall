package order

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidOrder = errs.New(60001, "invalid order")
	ErrInvalidMoney = errs.New(60002, "invalid money")
)
