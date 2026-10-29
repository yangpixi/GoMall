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
