package product

import (
	"errors"
	"strings"
)

type State struct {
	ID          int64
	Name        string
	ShopID      int64
	Description string
	Status      int
	SkuIDs      []int64
}

func Restore(s *State) (*Product, error) {
	if s == nil || strings.TrimSpace(s.Name) == "" || strings.TrimSpace(s.Description) == "" {
		return nil, errors.New("invalid product state")
	}

	return &Product{
		id:          s.ID,
		name:        s.Name,
		shopID:      s.ShopID,
		description: s.Description,
		status:      s.Status,
		skuIDs:      s.SkuIDs,
	}, nil
}
