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

func (s *SKU) Snapshot() (*State, error) {
	if s.id == 0 || s.productID == 0 || strings.TrimSpace(s.specification) == "" {
		return nil, ErrInvalidSKU
	}

	return &State{
		ID:            s.id,
		ProductID:     s.productID,
		Price:         s.price,
		Specification: s.specification,
	}, nil
}
