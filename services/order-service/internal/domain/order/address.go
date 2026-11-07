package order

import "strings"

// Address value object
type Address struct {
	address   string
	recipient string
	phone     string
}

func NewAddress(address, recipient, phone string) (*Address, error) {
	if strings.TrimSpace(address) == "" || strings.TrimSpace(recipient) == "" || strings.TrimSpace(phone) == "" {
		return nil, ErrInvalidAddress
	}

	return &Address{
		address:   address,
		recipient: recipient,
		phone:     phone,
	}, nil
}

type AddressState struct {
	Address   string
	Recipient string
	Phone     string
}

func RestoreAddress(s *AddressState) (*Address, error) {
	if strings.TrimSpace(s.Address) == "" || strings.TrimSpace(s.Recipient) == "" || strings.TrimSpace(s.Phone) == "" {
		return nil, ErrInvalidAddress
	}

	return &Address{
		address:   s.Address,
		recipient: s.Recipient,
		phone:     s.Phone,
	}, nil
}

func (a *Address) Snapshot() *AddressState {
	return &AddressState{
		Address:   a.address,
		Recipient: a.recipient,
		Phone:     a.phone,
	}
}
