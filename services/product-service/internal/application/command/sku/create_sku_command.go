package sku

type CreateSKUCommand struct {
	ProductID     int64
	Price         int64
	Specification string
}
