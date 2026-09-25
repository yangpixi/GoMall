package product

import "strings"

type Product struct {
	id          int64
	name        string
	description string
	status      int
	skuIDs      []int64
}

func New(id int64, name, des string, status int, skuIDs []int64) (*Product, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(des) == "" {
		return nil, ErrInvalidProduct
	}

	return &Product{
		id:          id,
		name:        name,
		description: des,
		status:      status,
		skuIDs:      skuIDs,
	}, nil
}
