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
