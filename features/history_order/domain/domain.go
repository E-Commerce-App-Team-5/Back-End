package domain

type Core struct {
	ID         uint
	IdProduct  uint
	ProductQty int
	Price      int
	NamaToko string
	ProductName string
	PriceSum int
	ProductPicture string
}

type Services interface {
	GetBuy(id uint) ([]Core, error)
	// GetSell(id uint) ([]Core, error)
}

type Repostory interface {
	GetBuy(id uint) ([]Core, error)
	// GetSell(id uint) ([]Core, error)
}
