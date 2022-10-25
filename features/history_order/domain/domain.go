package domain

type Core struct {
	ID          uint
	IdCheckout  uint
	IdProduct   uint
	ProductQty  int
	Price       int
	NamaToko string
	ProductName  string
	Status string
}

type CoreCheckout struct {
	ID          uint
	IdPembeli   uint
	OrderId     string
	GrossAmount float32
	Token       string
	Link        string
	Status      string
}

type CoreProduct struct {
	ID             uint
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductDetail  string
	ProductQty     int
	Price          int
	ProductPicture string
}

type Services interface {
	GetBuy(id uint) ([]Core, error)
	GetSell(id uint) ([]Core, error)
}

type Repostory interface {
	GetBuy(id uint) ([]Core, error)
	GetSell(id uint) ([]Core, error)
}
