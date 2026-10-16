package stock

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
)

type ReleaseHandler struct {
	repo reservation.Repository
}

func NewReleaseHandler(r reservation.Repository) (*ReleaseHandler, error) {
	if r == nil {
		return nil, errors.New("invalid release handler")
	}

	return &ReleaseHandler{repo: r}, nil
}

func (h *ReleaseHandler) Handle(ctx context.Context, cmd *ReleaseCommand) error {
	rev, err := h.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	rev.ToStatusReleased()

	err = h.repo.Update(ctx, rev)
	if err != nil {
		return err
	}

	return nil
}
