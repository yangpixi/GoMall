package stock

type ReserveCommand struct {
	Quantity int64
	SkuID    int64
	OrderID  int64
}
