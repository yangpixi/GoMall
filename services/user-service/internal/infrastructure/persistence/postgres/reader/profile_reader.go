package reader

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/yangpixi/GoMall/services/user-service/internal/application/query"
	"github.com/yangpixi/GoMall/services/user-service/internal/domain/profile"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ProfileReader struct {
	db *gorm.DB
}

func NewProfileReader(db *gorm.DB) *ProfileReader {
	return &ProfileReader{db: db}
}

func (r *ProfileReader) GetProfileDetail(ctx context.Context, userID int64) (*query.ProfileVO, error) {
	p, err := gorm.G[model.Profile](r.db).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to select profile: %d, error: %w", userID, err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, profile.ErrProfileNotFound
	}

	return &query.ProfileVO{
		UserID:   strconv.FormatInt(p.UserID, 10),
		Nickname: p.Nickname,
		Phone:    p.Phone,
		Email:    p.Email,
		Avatar:   p.Avatar,
	}, nil
}
