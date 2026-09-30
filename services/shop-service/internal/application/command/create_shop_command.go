package command

type CreateShopCommand struct {
	OwnerID     int64 // for admin
	Name        string
	Description string
}
