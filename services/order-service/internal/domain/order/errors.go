package order

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidOrder  = errs.New(60001, "invalid order")
	ErrInvalidMoney  = errs.New(60002, "invalid money")
	ErrOrderNotFound = errs.New(60003, "order not found")

	ErrInvalidAddress  = errs.New(61001, "invalid address")
	ErrAddressNotFound = errs.New(61002, "address not found")

	ErrInvalidItem  = errs.New(62001, "invalid item")
	ErrItemNotFound = errs.New(62002, "item not found")
)
