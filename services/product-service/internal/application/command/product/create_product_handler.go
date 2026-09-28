package product

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/product-service/internal/domain/product"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/id"
)

type CreateProductHandler struct {
	repo product.Repository
}

func NewCreateProductHandler(r product.Repository) (*CreateProductHandler, error) {
	if r == nil {
		return nil, errors.New("invalid product repository")
	}

	return &CreateProductHandler{repo: r}, nil
}

func (h *CreateProductHandler) Handle(ctx context.Context, cmd *CreateProductCommand) error {
	// TODO: refine id getting strategy, also for other services
	generator, err := id.NewGenerator(1)
	if err != nil {
		return err
	}

	// TODO: shopID should be retrieved from shop-service
	// shopID from command is used for admin endpoint
	p, err := product.New(generator.NextID(), cmd.ShopID, cmd.Name, cmd.Description, cmd.Status)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
