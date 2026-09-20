package mapper

import (
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
)

func ToAddress(a *model.UserAddress) (*address.Address, error) {
	if a == nil {
		return nil, errors.New("invalid address model")
	}

	return address.Restore(&address.State{
		ID:        a.ID,
		UserID:    a.UserID,
		Address:   a.Address,
		Phone:     a.Phone,
		Recipient: a.Recipient,
	})
}

func ToAddressPO(a *address.Address) (*model.UserAddress, error) {
	if a == nil {
		return nil, errors.New("invalid address entity")
	}

	s, err := a.Snapshot()
	if err != nil {
		return nil, errors.New("failed to take a snapshot of address entity")
	}

	return &model.UserAddress{
		ID:      s.ID,
		UserID:  s.UserID,
		Address: s.Address,
		Phone:   s.Phone,
	}, nil
}
