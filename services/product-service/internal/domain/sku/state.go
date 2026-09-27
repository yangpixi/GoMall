package sku

import "strings"

type State struct {
	ID            int64
	ProductID     int64
	Price         int
	Specification string
}

func Restore(s *State) (*SKU, error) {
	if s.ID == 0 || s.ProductID == 0 || strings.TrimSpace(s.Specification) == "" {
		return nil, ErrInvalidSKU
	}

	return &SKU{
		id:            s.ID,
		productID:     s.ProductID,
		price:         s.Price,
		specification: s.Specification,
	}, nil
}
