package command

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/account"
	"github.com/yangpixi/GoMall/auth-service/internal/domain/role"
	"github.com/yangpixi/GoMall/auth-service/internal/domain/shared"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	accountRepo account.Repository
	roleRepo    role.Repository
	idGenerator shared.IDGenerator
}

func NewRegisterHandler(accountRepo account.Repository, roleRepo role.Repository, generator shared.IDGenerator) (*RegisterHandler, error) {
	if accountRepo == nil || roleRepo == nil || generator == nil {
		return nil, errors.New("missing required arguments")
	}

	return &RegisterHandler{
		accountRepo: accountRepo,
		roleRepo:    roleRepo,
		idGenerator: generator,
	}, nil
}

func (h *RegisterHandler) Handle(ctx context.Context, cmd *RegisterCommand) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// by default, the role id 1 represents user
	a, err := account.New(uint(h.idGenerator.NextID()), cmd.Username, string(encryptedPassword), []uint{1})
	if err != nil {
		return err
	}

	err = h.accountRepo.Save(ctx, a)
	if err != nil {
		return err
	}

	err = h.roleRepo.GrantByUserID(ctx, a.ID(), a.RoleIDs())
	if err != nil {
		return err
	}

	return nil
}
