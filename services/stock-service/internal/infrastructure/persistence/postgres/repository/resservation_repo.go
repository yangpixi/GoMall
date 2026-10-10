package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ReservationRepo struct {
	db *gorm.DB
}

func NewReservationRepo(db *gorm.DB) *ReservationRepo {
	return &ReservationRepo{db: db}
}

func (r *ReservationRepo) FindByID(ctx context.Context, id int64) (*reservation.StockReservation, error) {
	sr, err := gorm.G[*model.StockReservation](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, reservation.ErrStockReservationNotFound
		}

		return nil, fmt.Errorf("failed to select stock reservation, id: %d, error: %w", id, err)
	}

	return mapper.ToStockReservation(sr)
}

func (r *ReservationRepo) FindByOrderID(ctx context.Context, orderID int64) (*reservation.StockReservation, error) {
	sr, err := gorm.G[*model.StockReservation](r.db).Where("order_id = ?", orderID).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, reservation.ErrStockReservationNotFound
		}

		return nil, fmt.Errorf("failed to select stock reservation,orderID: %d, error: %w", orderID, err)
	}

	return mapper.ToStockReservation(sr)
}

func (r *ReservationRepo) FindBySkuID(ctx context.Context, skuID int64) (*reservation.StockReservation, error) {
	sr, err := gorm.G[*model.StockReservation](r.db).Where("sku_id = ?", skuID).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, reservation.ErrStockReservationNotFound
		}

		return nil, fmt.Errorf("failed to select stock reservation, skuID: %d, error: %w", skuID, err)
	}

	return mapper.ToStockReservation(sr)
}

func (r *ReservationRepo) Save(ctx context.Context, sr *reservation.StockReservation) error {
	po, err := mapper.ToStockReservationPO(sr)
	if err != nil {
		return fmt.Errorf("failed to transfer to po: %w", err)
	}

	err = gorm.G[model.StockReservation](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save stock reservation: %w", err)
	}

	return nil
}
