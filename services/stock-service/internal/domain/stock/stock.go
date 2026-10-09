package stock

type Stock struct {
	skuID        int64
	availableQTY int64
	lockedQTY    int64
}

func New(skuID, availableQTY, lockedQty int64) (*Stock, error) {
	if skuID == 0 || availableQTY <= 0 || lockedQty <= 0 {
		return nil, ErrInvalidStock
	}

	return &Stock{
		skuID:        skuID,
		availableQTY: availableQTY,
		lockedQTY:    lockedQty,
	}, nil
}

func (s *Stock) Snapshot() (*State, error) {
	return &State{
		SkuID:        s.skuID,
		AvailableQTY: s.availableQTY,
		LockedQTY:    s.lockedQTY,
	}, nil
}
