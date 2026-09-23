package address

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type DeleteAddressHandler struct {
	repo address.Repository
}

func NewDeleteAddressHandler(repo address.Repository) (*DeleteAddressHandler, error) {
	if repo == nil {
		return nil, errors.New("invalid address repository")
	}

	return &DeleteAddressHandler{repo: repo}, nil
}

func (h *DeleteAddressHandler) Handle(ctx context.Context, cmd *DeleteAddressCommand) error {

	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return errors.New("failed to retrieve userID from context")
	}

	err := h.repo.Delete(ctx, cmd.ID, userID)
	if err != nil {
		return err
	}

	return nil
}
