package address

type Address struct {
	id        uint
	userID    uint
	address   string
	phone     string
	recipient string
}

func New(id, userID uint, address, phone, recipient string) (*Address, error) {
	if userID == 0 || address == "" || phone == "" || recipient == "" {
		return nil, ErrInvalidAddress
	}

	return &Address{
		id:        id,
		userID:    userID,
		address:   address,
		phone:     phone,
		recipient: recipient,
	}, nil
}
