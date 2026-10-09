package stock

type State struct {
	SkuID        int64
	AvailableQTY int64
	LockedQTY    int64
}

func Restore(s *State) (*Stock, error) {
	if s.SkuID == 0 || s.AvailableQTY <= 0 || s.LockedQTY <= 0 {
		return nil, ErrInvalidStock
	}

	return &Stock{
		skuID:        s.SkuID,
		availableQTY: s.AvailableQTY,
		lockedQTY:    s.LockedQTY,
	}, nil
}
