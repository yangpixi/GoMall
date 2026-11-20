package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) FindByID(ctx context.Context, id int64) (*payment.Payment, error) {
	mp, err := gorm.G[*model.Payment](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, payment.ErrPaymentNotFound
		}
		return nil, fmt.Errorf("failed to select payment, id: %d, error: %w", id, err)
	}

	p, err := mapper.ToPayment(mp)
	if err != nil {
		return nil, fmt.Errorf("failed to transform po")
	}

	return p, nil
}

func (r *PaymentRepo) Save(ctx context.Context, p *payment.Payment) error {
	po := mapper.ToPaymentPO(p)
	err := gorm.G[model.Payment](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save payment po")
	}

	return nil
}
