package product

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidProduct = errs.New(30001, "invalid product")
)
