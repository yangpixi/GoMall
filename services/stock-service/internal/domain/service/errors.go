package service

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidSKUID = errs.New(52001, "invalid sku id")
)
