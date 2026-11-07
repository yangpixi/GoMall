package order

type State struct {
	ID             int64
	UserID         int64
	Status         Status
	TotalAmount    int64
	DiscountAmount int64
	FinalAmount    int64

	Items   []*Item
	Address *Address
}

func Restore(s *State) (*Order, error) {
	if s.ID == 0 || s.UserID == 0 {
		return nil, ErrInvalidOrder
	}

	if len(s.Items) == 0 {
		return nil, ErrInvalidOrder
	}

	if s.Address == nil {
		return nil, ErrInvalidOrder
	}

	if s.TotalAmount < 0 || s.DiscountAmount < 0 || s.FinalAmount < 0 {
		return nil, ErrInvalidMoney
	}

	if s.TotalAmount != s.DiscountAmount+s.FinalAmount {
		return nil, ErrInvalidMoney
	}

	return &Order{
		id:             s.ID,
		userID:         s.UserID,
		status:         s.Status,
		totalAmount:    s.TotalAmount,
		discountAmount: s.DiscountAmount,
		finalAmount:    s.FinalAmount,
		items:          s.Items,
		address:        s.Address,
	}, nil
}
