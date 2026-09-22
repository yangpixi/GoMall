package address

import (
	"regexp"
	"strings"
)

type Address struct {
	id        int64
	userID    int64
	address   string
	phone     string
	recipient string
}

var phonePattern = regexp.MustCompile(`^1[3-9][0-9]{9}$`)

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

func (a *Address) Snapshot() (*State, error) {
	return &State{
		ID:        a.id,
		UserID:    a.userID,
		Address:   a.address,
		Phone:     a.phone,
		Recipient: a.recipient,
	}, nil
}

func (a *Address) ChangeAddress(address string) error {
	if strings.TrimSpace(address) == "" {
		return ErrInvalidAddress
	}

	a.address = address
	return nil
}

func (a *Address) ChangePhone(phone string) error {
	phone = strings.TrimSpace(phone)

	if !phonePattern.MatchString(phone) {
		return ErrInvalidPhone
	}

	a.phone = phone
	return nil
}

func (a *Address) ChangeRecipient(r string) error {
	if strings.TrimSpace(r) == "" {
		return ErrInvalidRecipient
	}

	a.recipient = r
	return nil
}
