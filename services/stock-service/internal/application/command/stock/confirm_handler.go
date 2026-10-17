package stock

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
)

type ConfirmHandler struct {
	repo reservation.Repository
}

func NewConfirmHandler(r reservation.Repository) (*ConfirmHandler, error) {
	if r == nil {
		return nil, errors.New("invalid release handler")
	}

	return &ConfirmHandler{repo: r}, nil
}

func (h *ConfirmHandler) Handle(ctx context.Context, cmd *ConfirmCommand) error {
	rev, err := h.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	rev.ToStatusConfirmed()

	err = h.repo.Update(ctx, rev)
	if err != nil {
		return err
	}

	return nil
}
