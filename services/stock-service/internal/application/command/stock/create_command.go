package stock

type CreateSKUStockCommand struct {
	SkuID        int64
	AvailableQTY int64
	LockedQTY    int64
}
