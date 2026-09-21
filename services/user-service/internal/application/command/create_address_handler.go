package command

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
	"github.com/yangpixi/GoMall/services/user-service/internal/domain/shared"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreateAddressHandler struct {
	repo        address.Repository
	idGenerator shared.IDGenerator
}

func NewCreateAddressHandler(repo address.Repository, id shared.IDGenerator) (*CreateAddressHandler, error) {
	return &CreateAddressHandler{
		repo:        repo,
		idGenerator: id,
	}, nil
}

func (h *CreateAddressHandler) Handle(ctx context.Context, cmd *CreateAddressCommand) error {
	snowflakeID := h.idGenerator.NextID()
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return errors.New("failed to retrieve id from context")
	}

	add, err := address.New(snowflakeID, userID, cmd.Address, cmd.Phone, cmd.Recipient)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, add)
	if err != nil {
		return err
	}

	return nil
}
