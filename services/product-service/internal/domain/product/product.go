package product

import "strings"

type Product struct {
	id          int64
	name        string
	shopID      int64
	description string
	status      int
	skuIDs      []int64
}

func New(id, shopID int64, name, des string, status int, skuIDs []int64) (*Product, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(des) == "" || id == 0 || shopID == 0 {
		return nil, ErrInvalidProduct
	}

	return &Product{
		id:          id,
		name:        name,
		shopID:      shopID,
		description: des,
		status:      status,
		skuIDs:      skuIDs,
	}, nil
}

func (p *Product) Snapshot() (*State, error) {
	if p.id == 0 || p.shopID == 0 || strings.TrimSpace(p.name) == "" || strings.TrimSpace(p.description) == "" {
		return nil, ErrInvalidProduct
	}

	return &State{
		ID:          p.id,
		Name:        p.name,
		ShopID:      p.shopID,
		Description: p.description,
		Status:      p.status,
		SkuIDs:      p.skuIDs,
	}, nil
}
