package product

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidProduct  = errs.New(30001, "invalid product")
	ErrProductNotFound = errs.New(30002, "product not found")
)
