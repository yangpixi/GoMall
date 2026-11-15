package payment

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidPayment  = errs.New(70001, "invalid payment")
	ErrPaymentNotFound = errs.New(70002, "payment not found")
)
