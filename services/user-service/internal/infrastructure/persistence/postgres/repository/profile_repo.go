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

func (r *ProfileRepo) FindByUserID(ctx context.Context, userID int64) (*profile.Profile, error) {
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

	addIDs := make([]int64, 0)
	err = r.db.WithContext(ctx).Model(&model.UserAddress{}).Where("user_id = ?", userID).Pluck("id", &addIDs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query user %d's profile %w", userID, err)
	}

	return mapper.ToProfile(p, addIDs)
}

func (r *ProfileRepo) Save(ctx context.Context, p *profile.Profile) error {
	po, err := mapper.ToProfilePO(p)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	err = gorm.G[model.Profile](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	return nil
}

func (r *ProfileRepo) Update(ctx context.Context, p *profile.Profile) error {
	po, err := mapper.ToProfilePO(p)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	rowsAffected, err := gorm.G[map[string]any](r.db).Table("profile").Where("user_id = ?", po.UserID).Updates(ctx, map[string]any{
		"nickname": po.Nickname,
		"phone":    po.Phone,
		"email":    po.Email,
	})

	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	if rowsAffected == 0 {
		return profile.ErrProfileNotFound
	}

	return nil
}
