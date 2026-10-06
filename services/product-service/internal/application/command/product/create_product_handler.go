package product

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/product-service/internal/domain/product"
	"github.com/yangpixi/GoMall/services/product-service/internal/domain/shared"
)

type CreateProductHandler struct {
	repo        product.Repository
	idGenerator shared.IDGenerator
}

func NewCreateProductHandler(r product.Repository, generator shared.IDGenerator) (*CreateProductHandler, error) {
	if r == nil || generator == nil {
		return nil, errors.New("invalid product repository")
	}

	return &CreateProductHandler{repo: r, idGenerator: generator}, nil
}

func (h *CreateProductHandler) Handle(ctx context.Context, cmd *CreateProductCommand) error {
	// TODO: shopID should be retrieved from shop-service
	// shopID from command is used for admin endpoint
	p, err := product.New(h.idGenerator.NextID(), cmd.ShopID, cmd.Name, cmd.Description, cmd.Status)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
