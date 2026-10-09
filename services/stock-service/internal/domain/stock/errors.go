package stock

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidStock     = errs.New(50001, "invalid stock")
	ErrSkuStockNotFound = errs.New(50002, "sku stock not found")
)
