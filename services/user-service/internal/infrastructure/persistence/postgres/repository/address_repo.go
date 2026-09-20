package repository

import (
	"context"
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

func (r *AddressRepo) FindByUserID(ctx context.Context, userID int64) ([]*address.Address, error) {
	// len(addresses) might be zero
	addresses, err := gorm.G[*model.UserAddress](r.db).Where("user_id = ?", userID).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to select %d's address: %w", userID, err)
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
