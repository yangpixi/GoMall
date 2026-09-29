package shop

import "strings"

type State struct {
	ID          int64
	OwnerID     int64
	Name        string
	Description string
	IsBanned    int
}

func Restore(s *State) (*Shop, error) {
	if s.ID == 0 || s.OwnerID == 0 || strings.TrimSpace(s.Name) == "" || strings.TrimSpace(s.Description) == "" {
		return nil, ErrInvalidShop
	}

	return &Shop{
		id:          s.ID,
		ownerID:     s.OwnerID,
		name:        s.Name,
		description: s.Description,
		isBanned:    s.IsBanned,
	}, nil
}
