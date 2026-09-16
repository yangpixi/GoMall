package address

import "errors"

type State struct {
	ID        uint
	UserID    uint
	Address   string
	Phone     string
	Recipient string
}

func Restore(s *State) (*Address, error) {
	if s.UserID == 0 || s.Address == "" || s.Phone == "" || s.Recipient == "" {
		return nil, errors.New("invalid arguments")
	}

	return &Address{
		id:        s.ID,
		userID:    s.UserID,
		address:   s.Address,
		phone:     s.Phone,
		recipient: s.Recipient,
	}, nil
}
