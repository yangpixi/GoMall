package mapper

import (
	"fmt"

	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/model"
)

func ToItem(item *model.GomallOrderItem) (*order.Item, error) {
	return order.RestoreItem(&order.ItemStatus{
		ID:         item.ID,
		SkuID:      item.SkuID,
		Quantity:   item.Quantity,
		UnitPrice:  item.UnitPrice,
		TotalPrice: item.TotalPrice,
	})
}

func ToAddress(a *model.GomallOrderAddress) (*order.Address, error) {
	return order.RestoreAddress(&order.AddressState{
		Address:   a.Address,
		Recipient: a.Recipient,
		Phone:     a.Phone,
	})
}

func ToOrder(o *model.GomallOrder, items []*model.GomallOrderItem, a *model.GomallOrderAddress) (*order.Order, error) {
	is := make([]*order.Item, 0, len(items))
	for _, item := range items {
		i, err := ToItem(item)
		if err != nil {
			return nil, fmt.Errorf("failed to transfer item po")
		}
		is = append(is, i)
	}

	address, err := ToAddress(a)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer address po")
	}

	return order.Restore(&order.State{
		ID:             o.ID,
		UserID:         o.UserID,
		Status:         o.Status,
		TotalAmount:    o.TotalAmount,
		DiscountAmount: o.DiscountAmount,
		FinalAmount:    o.FinalAmount,
		Items:          is,
		Address:        address,
	})
}

// to po

func ToOrderPO(o *order.Order) *model.GomallOrder {
	snap := o.Snapshot()
	return &model.GomallOrder{
		ID:             snap.ID,
		UserID:         snap.UserID,
		Status:         snap.Status,
		TotalAmount:    snap.TotalAmount,
		DiscountAmount: snap.DiscountAmount,
		FinalAmount:    snap.FinalAmount,
	}
}

func ToItemPO(i *order.Item, orderID int64) *model.GomallOrderItem {
	snap := i.Snapshot()
	return &model.GomallOrderItem{
		ID:         snap.ID,
		OrderID:    orderID,
		SkuID:      snap.SkuID,
		Quantity:   snap.Quantity,
		UnitPrice:  snap.UnitPrice,
		TotalPrice: snap.TotalPrice,
	}
}

func ToAddressPO(add *order.Address, orderID int64) *model.GomallOrderAddress {
	snap := add.Snapshot()
	return &model.GomallOrderAddress{
		OrderID:   orderID,
		Address:   snap.Address,
		Recipient: snap.Recipient,
		Phone:     snap.Phone,
	}
}
