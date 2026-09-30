package sku

import (
	"context"

	"github.com/yangpixi/GoMall/product-service/internal/domain/sku"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/id"
)

type CreateSKUHandler struct {
	repo sku.Repository
}

func NewCreateSKUHandler(repo sku.Repository) (*CreateSKUHandler, error) {
	return &CreateSKUHandler{repo: repo}, nil
}

func (h *CreateSKUHandler) Handle(ctx context.Context, cmd *CreateSKUCommand) error {
	generator, err := id.NewGenerator(1)
	if err != nil {
		return err
	}

	s, err := sku.New(generator.NextID(), cmd.ProductID, cmd.Price, cmd.Specification)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
