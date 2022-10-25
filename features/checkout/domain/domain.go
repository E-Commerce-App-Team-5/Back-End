package domain

type Core struct {
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

type HistoryCore struct {
	ID         uint
	IdPembeli  uint
	IdProduct  uint
	ProductQty int
	Price      int
}

type Repository interface {
	Insert(newHistory []HistoryCore, newCheckout Core) (Core, error)
	Get(id uint) ([]Core, error)
	Delete(id uint) (Core, error)
}

type Service interface { // Bisnis logic
	AddCheckout(newHistory []HistoryCore, newCheckout Core) (Core, error)
	DeleteCheckout(ID uint) (Core, error)
	GetCheckout(id uint) ([]Core, error)
}
