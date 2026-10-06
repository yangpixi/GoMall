package command

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/shop-service/internal/domain/shop"
	uID "github.com/yangpixi/GoMall/shared/http/id"
)

type CreateShopHandler struct {
	repo        shop.Repository
	idGenerator shop.IDGenerator
}

func NewCreateShopHandler(repo shop.Repository, generator shop.IDGenerator) (*CreateShopHandler, error) {
	if repo == nil || generator == nil {
		return nil, errors.New("invalid shop repository")
	}

	return &CreateShopHandler{repo: repo, idGenerator: generator}, nil
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
