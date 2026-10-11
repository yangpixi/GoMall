package command

import (
	"context"
	"errors"
	"time"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/service"
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/shared"
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/stock"
)

type ReserveHandler struct {
	stockRepo       stock.Repository
	reservationRepo reservation.Repository
	idGenerator     shared.IDGenerator
}

func NewReserveHandler(sr stock.Repository, rr reservation.Repository, generator shared.IDGenerator) (*ReserveHandler, error) {
	if sr == nil || rr == nil || generator == nil {
		return nil, errors.New("invalid reserve handler ")
	}

	return &ReserveHandler{
		stockRepo:       sr,
		reservationRepo: rr,
		idGenerator:     generator,
	}, nil
}

func (h *ReserveHandler) Handle(ctx context.Context, cmd *ReserveCommand) error {
	sku, err := h.stockRepo.FindBySkuID(ctx, cmd.SkuID)
	if err != nil {
		return err
	}

	sr, err := reservation.New(h.idGenerator.NextID(), cmd.OrderID, cmd.SkuID, cmd.Quantity, time.Now().Add(30*time.Minute))
	if err != nil {
		return err
	}

	err = service.Reserve(sku, sr)
	if err != nil {
		return err
	}

	return nil
}
