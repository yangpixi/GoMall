package mapper

import (
	"errors"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/account"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/model"
)

func ToAccount(p *model.User, roleIDs []uint) (*account.Account, error) {
	if p == nil {
		return nil, errors.New("user object is nil")
	}

	if p.Username == "" || p.Password == "" {
		return nil, errors.New("invalid user object")
	}

	return account.Restore(&account.State{
		ID:       p.ID,
		Username: p.Username,
		Password: p.Password,
		Status:   p.Status,
		RoleIDs:  roleIDs,
	})
}
func ToAccountPO(account *account.Account) (*model.User, error) {
	s, err := account.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       s.ID,
		Username: s.Username,
		Password: s.Password,
		Status:   s.Status,
	}, nil
}
