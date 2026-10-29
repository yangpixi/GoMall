package order

type Order struct {
	id             int64
	userID         int64
	status         Status
	totalAmount    int64
	discountAmount int64
	finalAmount    int64

	items   []*Item
	address *Address
}

func NewOrder(
	id, userID, totalAmount, discountAmount, finalAmount int64,
	items []*Item,
	address *Address,
) (*Order, error) {
	if id == 0 || userID == 0 {
		return nil, ErrInvalidOrder
	}

	if len(items) == 0 {
		return nil, ErrInvalidOrder
	}

	if address == nil {
		return nil, ErrInvalidOrder
	}

	if totalAmount < 0 || discountAmount < 0 || finalAmount < 0 {
		return nil, ErrInvalidMoney
	}

	if totalAmount != discountAmount+finalAmount {
		return nil, ErrInvalidMoney
	}

	return &Order{
		id:             id,
		userID:         userID,
		status:         StatusPendingPayment,
		totalAmount:    totalAmount,
		discountAmount: discountAmount,
		finalAmount:    finalAmount,
		items:          items,
		address:        address,
	}, nil

}

func (o *Order) ToStatusPaid() {
	o.status = StatusPaid
}

func (o *Order) ToStatusCancelled() {
	o.status = StatusCancelled
}

func (o *Order) ToStatusPendingPayment() {
	o.status = StatusPendingPayment
}
