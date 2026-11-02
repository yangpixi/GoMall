package order

type Item struct {
	id       int64
	skuID    int64
	quantity int64

	unitPrice  int64
	totalPrice int64
}

func NewItem(id, skuID, quantity, unitPrice, totalPrice int64) (*Item, error) {
	if id == 0 || skuID == 0 || quantity < 0 || unitPrice < 0 || totalPrice < 0 {
		return nil, ErrInvalidItem
	}

	return &Item{
		id:         id,
		skuID:      skuID,
		quantity:   quantity,
		unitPrice:  unitPrice,
		totalPrice: totalPrice,
	}, nil
}

type ItemStatus struct {
	ID       int64
	SkuID    int64
	Quantity int64

	UnitPrice  int64
	TotalPrice int64
}

func RestoreItem(s *ItemStatus) (*Item, error) {
	if s.ID == 0 || s.SkuID == 0 || s.Quantity < 0 || s.UnitPrice < 0 || s.TotalPrice < 0 {
		return nil, ErrInvalidItem
	}

	return &Item{
		id:         s.ID,
		skuID:      s.SkuID,
		quantity:   s.Quantity,
		unitPrice:  s.UnitPrice,
		totalPrice: s.TotalPrice,
	}, nil
}

func (i *Item) Snapshot() *ItemStatus {
	return &ItemStatus{
		ID:         i.id,
		SkuID:      i.skuID,
		Quantity:   i.quantity,
		UnitPrice:  i.unitPrice,
		TotalPrice: i.totalPrice,
	}
}
