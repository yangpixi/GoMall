package address

type DeleteAddressCommand struct {
	UserID int64 // only for admin side
	ID     int64 // address primary key
}
