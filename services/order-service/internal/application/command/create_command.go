package command

type CreateOrderCommand struct {
	Items   []*ItemCommand
	Address AddressCommand
}

type ItemCommand struct {
	SkuID    int64
	Quantity int64
}

type AddressCommand struct {
	Address   string
	Recipient string
	Phone     string
}
