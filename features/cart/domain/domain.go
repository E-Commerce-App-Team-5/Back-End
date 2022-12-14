package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	IdProduct      uint
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductQty     int
	ProductDetail  string
	Price          int
	ProductPicture string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newProduct Core) (Core, error)
	Get(id uint) ([]Core, error)
	Delete(id uint) (Core, error)
	Edit(input Core) (Core, error)
}

type Service interface { // Bisnis logic
	AddCart(newProduct Core) (Core, error)
	DeleteCart(ID uint) (Core, error)
	GetCart(id uint) ([]Core, error)
	UpdateCart(input Core) (Core, error)
}

type Handler interface {
	AddCart() echo.HandlerFunc
	GetCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}
