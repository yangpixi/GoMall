package reservation

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidStockReservation  = errs.New(51001, "invalid stock reservation")
	ErrStockReservationNotFound = errs.New(51002, "stock reservation not found")
)
