package address

type UpdateAddressCommand struct {
	ID        int64
	Address   *string
	Phone     *string
	Recipient *string
}
