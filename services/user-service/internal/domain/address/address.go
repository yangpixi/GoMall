package address

type Address struct {
	id        int64
	userID    int64
	address   string
	phone     string
	recipient string
}

func New(id, userID int64, address, phone, recipient string) (*Address, error) {
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
