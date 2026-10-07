package product

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/product-service/internal/application/port"
	"github.com/yangpixi/GoMall/services/product-service/internal/domain/product"
	"github.com/yangpixi/GoMall/services/product-service/internal/domain/shared"
)

type CreateProductHandler struct {
	repo        product.Repository
	idGenerator shared.IDGenerator
	shopReader  port.ShopIDReader
}

func NewCreateProductHandler(r product.Repository, generator shared.IDGenerator, sReader port.ShopIDReader) (*CreateProductHandler, error) {
	if r == nil || generator == nil || sReader == nil {
		return nil, errors.New("invalid product repository")
	}

	return &CreateProductHandler{repo: r, idGenerator: generator, shopReader: sReader}, nil
}

func (h *CreateProductHandler) Handle(ctx context.Context, cmd *CreateProductCommand) error {
	sInfo, err := h.shopReader.GetShopID(ctx)
	if err != nil {
		return err
	}

	p, err := product.New(h.idGenerator.NextID(), sInfo.ShopID, cmd.Name, cmd.Description, cmd.Status)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
