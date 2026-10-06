package mapper

import (
	"github.com/yangpixi/GoMall/services/shop-service/internal/domain/shop"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/model"
)

func ToShop(s *model.Shop) (*shop.Shop, error) {
	return shop.Restore(&shop.State{
		ID:          s.ID,
		OwnerID:     s.OwnerID,
		Name:        s.Name,
		Description: s.Description,
		IsBanned:    s.IsBanned,
	})
}

func ToShopPO(s *shop.Shop) (*model.Shop, error) {
	ss, err := s.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.Shop{
		ID:          ss.ID,
		OwnerID:     ss.OwnerID,
		Name:        ss.Name,
		Description: ss.Description,
		IsBanned:    ss.IsBanned,
	}, nil
}
