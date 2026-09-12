package repository

import (
	"context"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/account"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) account.Repository {
	return &AccountRepo{db: db}
}

func (a *AccountRepo) FindByID(ctx context.Context, id uint) (*account.Account, error) {
	po, err := gorm.G[model.User](a.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	var ids []uint

	err = a.db.WithContext(ctx).Model(&model.UserRole{}).Where("user_id = ?", id).Pluck("role_id", &ids).Error
	if err != nil {
		return nil, err
	}

	acc, err := mapper.ToAccount(&po, ids)
	return acc, err
}

func (a *AccountRepo) FindByUsername(ctx context.Context, username string) (*account.Account, error) {
	po, err := gorm.G[model.User](a.db).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, account.ErrAccountNotFound
	}

	var ids []uint

	err = a.db.WithContext(ctx).Model(&model.UserRole{}).Where("user_id = ?", po.ID).Pluck("role_id", &ids).Error
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, account.ErrAccountNotFound
	}

	acc, err := mapper.ToAccount(&po, ids)
	return acc, err
}

func (a *AccountRepo) Save(ctx context.Context, account *account.Account) error {
	po, err := mapper.ToAccountPO(account)
	if err != nil {
		return err
	}

	err = gorm.G[model.User](a.db).Create(ctx, po)
	if err != nil {
		return err
	}

	return nil
}
