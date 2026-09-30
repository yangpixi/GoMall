package shop

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidShop  = errs.New(40001, "invalid shop profile")
	ErrShopNotFound = errs.New(40002, "shop not found")
)
