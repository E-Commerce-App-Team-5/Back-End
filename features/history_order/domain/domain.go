package domain

type Core struct {
	ID         uint
	IdProduct  uint
	ProductQty int
	Price      int
	NamaToko string
	NamaPembeli string
	ProductName string
	ProductsDetail string
	PriceSum int
	ProductPicture string
}

type Services interface {
	GetBuy(id uint) ([]Core, error)
	GetSell(id uint) ([]Core, error)
}

type Repository interface {
	GetBuy(id uint) ([]Core, error)
	GetSell(id uint) ([]Core, error)
}
