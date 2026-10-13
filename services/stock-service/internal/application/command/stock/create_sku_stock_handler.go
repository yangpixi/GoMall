package stock

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/stock"
)

type CreateSKUStockHandler struct {
	repo stock.Repository
}

func NewCreateSKUStockHandler(r stock.Repository) (*CreateSKUStockHandler, error) {
	if r == nil {
		return nil, errors.New("invalid create sku stock handler")
	}

	return &CreateSKUStockHandler{repo: r}, nil
}

func (h *CreateSKUStockHandler) Handle(ctx context.Context, cmd *CreateSKUStockCommand) error {
	s, err := stock.New(cmd.SkuID, cmd.AvailableQTY, cmd.LockedQTY)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, s)
	if err != nil {
		return err
	}

	return nil
}
