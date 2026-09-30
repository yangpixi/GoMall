package command

import (
	"context"
	"errors"

	uID "github.com/yangpixi/GoMall/shared/http/id"
	"github.com/yangpixi/GoMall/shop-service/internal/domain/shop"
)

type CreateShopHandler struct {
	repo        shop.Repository
	idGenerator shop.IDGenerator
}

func NewCreateShopHandler(repo shop.Repository) (*CreateShopHandler, error) {
	if repo == nil {
		return nil, errors.New("invalid shop repository")
	}

	return &CreateShopHandler{repo: repo}, nil
}

func (h *CreateShopHandler) Handle(ctx context.Context, cmd *CreateShopCommand) error {
	userID, ok := uID.UserIDFromCtx(ctx)
	if !ok {
		return errors.New("failed to retrieve userID from context")
	}

	id := h.idGenerator.NextID()

	s, err := shop.New(id, userID, cmd.Name, cmd.Description)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
