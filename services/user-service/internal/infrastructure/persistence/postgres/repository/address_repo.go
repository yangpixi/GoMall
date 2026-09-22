package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/address"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type AddressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) address.Repository {
	return &AddressRepo{db: db}
}

func (r *AddressRepo) FindByID(ctx context.Context, id int64) (*address.Address, error) {
	a, err := gorm.G[*model.UserAddress](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, address.ErrAddressNotFound
		}
		return nil, fmt.Errorf("failed to query address, id: %d, error: %w", id, err)
	}

	return mapper.ToAddress(a)
}

func (r *AddressRepo) FindByUserID(ctx context.Context, userID int64) ([]*address.Address, error) {
	// len(addresses) might be zero
	addresses, err := gorm.G[*model.UserAddress](r.db).Where("user_id = ?", userID).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query %d's address: %w", userID, err)
	}

	res := make([]*address.Address, 0, len(addresses))
	for _, add := range addresses {
		a, err := mapper.ToAddress(add)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}

	return res, nil
}

func (r *AddressRepo) Save(ctx context.Context, a *address.Address) error {
	po, err := mapper.ToAddressPO(a)
	if err != nil {
		return err
	}

	err = gorm.G[model.UserAddress](r.db).Create(ctx, po)
	if err != nil {
		return err
	}

	return nil
}

func (r *AddressRepo) Delete(ctx context.Context, id, userID int64) error {
	rowsAffected, err := gorm.G[model.UserAddress](r.db).Where("id = ? AND user_id = ?", id, userID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete address: %w", err)
	}

	if rowsAffected == 0 {
		return address.ErrAddressNotFound
	}

	return nil
}
