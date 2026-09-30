package product

type CreateProductCommand struct {
	Name        string
	ShopID      int64
	Description string
	Status      int
}
