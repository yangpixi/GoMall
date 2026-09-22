package address

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
)

type UpdateAddressHandler struct {
	repo address.Repository
}

func NewUpdateAddressHandler(repo address.Repository) (*UpdateAddressHandler, error) {
	if repo == nil {
		return nil, errors.New("invalid address repository")
	}

	return &UpdateAddressHandler{repo: repo}, nil
}

func (h *UpdateAddressHandler) Handle(ctx context.Context, cmd *UpdateAddressCommand) error {

	a, err := h.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if cmd.Phone != nil {
		if err := a.ChangePhone(*cmd.Phone); err != nil {
			return err
		}
	}

	if cmd.Address != nil {
		if err := a.ChangeAddress(*cmd.Address); err != nil {
			return err
		}
	}

	if cmd.Recipient != nil {
		if err := a.ChangeRecipient(*cmd.Recipient); err != nil {
			return err
		}
	}

	return nil
}
