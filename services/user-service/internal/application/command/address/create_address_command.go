package address

type CreateAddressCommand struct {
	UserID    int64 // only for admin side
	Address   string
	Phone     string
	Recipient string
}
