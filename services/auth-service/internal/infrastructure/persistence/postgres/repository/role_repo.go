package repository

import (
	"context"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/role"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) role.Repository {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) FindByUserID(ctx context.Context, id int64) ([]*role.Role, error) {
	userRole, err := gorm.G[model.UserRole](r.db).Where("user_id = ?", id).Find(ctx)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(userRole))
	for _, ur := range userRole {
		ids = append(ids, ur.RoleID)
	}

	rolePOs, err := gorm.G[model.Role](r.db).Where("id IN ?", ids).Find(ctx)
	if err != nil {
		return nil, err
	}

	roles := make([]*role.Role, 0, len(rolePOs))
	for _, rp := range rolePOs {
		pIDs := make([]int64, 0)
		err = r.db.WithContext(ctx).Model(&model.RolePermission{}).Where("role_id IN ?", &ids).Pluck("permission_id", &pIDs).Error
		if err != nil {
			return nil, err
		}
		if len(pIDs) == 0 {
			return nil, role.ErrPermissionNotFound
		}

		re, err := mapper.ToRole(&rp, pIDs)
		if err != nil {
			return nil, err
		}

		roles = append(roles, re)
	}

	return roles, nil
}

func (r *RoleRepo) GrantByUserID(ctx context.Context, userID int64, roleIDs []int64) error {

	urs := make([]model.UserRole, 0, len(roleIDs))

	for _, id := range roleIDs {
		ur := model.UserRole{
			UserID: userID,
			RoleID: id,
		}
		urs = append(urs, ur)
	}

	err := gorm.G[model.UserRole](r.db).CreateInBatches(ctx, &urs, len(roleIDs))
	if err != nil {
		return err
	}

	return nil
}
