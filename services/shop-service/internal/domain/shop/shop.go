package shop

import "strings"

type Shop struct {
	id          int64
	ownerID     int64
	name        string
	description string
	isBanned    int
}

func New(id, ownerID int64, name, description string) (*Shop, error) {
	if id == 0 || ownerID == 0 || strings.TrimSpace(name) == "" || strings.TrimSpace(description) == "" {
		return nil, ErrInvalidShop
	}

	return &Shop{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		isBanned:    0,
	}, nil
}

func (s *Shop) Snapshot() (*State, error) {
	return &State{
		ID:          s.id,
		OwnerID:     s.ownerID,
		Name:        s.name,
		Description: s.description,
		IsBanned:    s.isBanned,
	}, nil
}

func (s *Shop) IsBanned() bool {
	return s.isBanned == 1
}
