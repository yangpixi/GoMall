package model

import "time"

type GomallOrderItem struct {
	ID         int64 `gorm:"primaryKey"`
	OrderID    int64
	SkuID      int64
	Quantity   int64
	UnitPrice  int64
	TotalPrice int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
