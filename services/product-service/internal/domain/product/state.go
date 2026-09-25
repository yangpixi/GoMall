package product

import (
	"errors"
	"strings"
)

type State struct {
	ID          int64
	Name        string
	Description string
	Status      int
	SkuIDs      []int64
}

func Snapshot(s *State) (*Product, error) {
	if s == nil || strings.TrimSpace(s.Name) == "" || strings.TrimSpace(s.Description) == "" {
		return nil, errors.New("invalid product state")
	}

	return &Product{
		id:          s.ID,
		name:        s.Name,
		description: s.Description,
		status:      s.Status,
		skuIDs:      s.SkuIDs,
	}, nil
}
