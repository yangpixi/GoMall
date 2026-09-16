package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfileRepo(db *gorm.DB) profile.Repository {
	return &ProfileRepo{db: db}
}

func (r *ProfileRepo) FindByUserID(ctx context.Context, userID uint) (*profile.Profile, error) {
	if userID == 0 {
		return nil, profile.ErrInvalidUserID
	}

	p, err := gorm.G[*model.Profile](r.db).Where("user_id = ?", userID).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profile.ErrProfileNotFound
		}
		return nil, fmt.Errorf("failed to query user %d's profile %w", userID, err)
	}

	addIDs := make([]uint, 0)
	err = r.db.WithContext(ctx).Model(&model.UserAddress{}).Where("user_id = ?", userID).Pluck("id", &addIDs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query user %d's profile %w", userID, err)
	}

	return mapper.ToProfile(p, addIDs)
}
