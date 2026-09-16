package mapper

import (
	"errors"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/role"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/model"
)

func ToRole(po *model.Role, permissionIDs []uint) (*role.Role, error) {
	if po == nil {
		return nil, errors.New("invalid role model")
	}

	return role.Restore(&role.State{
		ID:            po.ID,
		Code:          po.Code,
		Name:          po.Name,
		PermissionIDs: permissionIDs,
	})
}
