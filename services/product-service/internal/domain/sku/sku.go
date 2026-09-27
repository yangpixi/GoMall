package sku

import "strings"

type SKU struct {
	id            int64
	productID     int64
	price         int
	specification string
}

func New(id, productID int64, price int, spec string) (*SKU, error) {
	if id == 0 || productID == 0 || strings.TrimSpace(spec) == "" {
		return nil, ErrInvalidSKU
	}

	return &SKU{
		id:            id,
		productID:     productID,
		price:         price,
		specification: spec,
	}, nil
}
