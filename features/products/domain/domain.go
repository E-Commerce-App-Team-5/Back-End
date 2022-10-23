package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID             uint
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductDetail  string
	ProductQty     int
	Price          int
	ProductPicture string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newProduct Core) (Core, error)
	Get(page int) ([]Core, error)
	Delete(id uint) (Core, error)
	Edit(input Core) (Core, error)
}

type Service interface { // Bisnis logic
	AddProduct(newProduct Core) (Core, error)
	DeleteProduct(ID uint) (Core, error)
	UpdateProduct(input Core) (Core, error)
	GetProduct(page int) ([]Core, error)
}

type Handler interface {
	AddProduct() echo.HandlerFunc
	GetProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}
