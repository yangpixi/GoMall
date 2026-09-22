package address

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidAddress   = errs.New(21001, "invalid address")
	ErrInvalidPhone     = errs.New(21002, "invalid phone number")
	ErrInvalidRecipient = errs.New(21003, "invalid recipient")
	ErrAddressNotFound  = errs.New(21004, "address not found")
)
