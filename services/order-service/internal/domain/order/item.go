package order

type Item struct {
	id       int64
	skuID    int64
	quantity int64

	unitPrice  int64
	totalPrice int64
}
