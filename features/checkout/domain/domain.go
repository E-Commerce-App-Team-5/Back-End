package domain

type Core struct {
	ID        uint
	IdPembeli uint
	Token     string
	Link      string
	Status    string
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

type DetailCore struct {
	ID         uint
	IdPembeli  uint
	IdProduct  uint
	ProductQty int
	Price      int
	Status     string
}

type Repository interface {
	Insert(newCheckout DetailCore) (Core, error)
	Get() ([]Core, error)
	Delete(id uint) (Core, error)
}

type Service interface { // Bisnis logic
	AddCheckout(newProduct Core) (Core, error)
	DeleteCheckout(ID uint) (Core, error)
	GetCheckout() ([]Core, error)
}
