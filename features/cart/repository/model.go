package repository

import (
	"ecommerce/features/cart/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname    string
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Password    string
	UserPicture string
	DOB         string
	Phone       string
	NamaToko    string
	Products    []Product `gorm:"foreignKey:IdUser"`
	Carts       []Cart    `gorm:"foreignKey:IdUser"`
}

type Product struct {
	gorm.Model
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductDetail  string
	ProductQty     int
	Price          int
	ProductPicture string
	Carts          []Cart `gorm:"foreignKey:IdProduct"`
}

type Cart struct {
	gorm.Model
	IdProduct      uint
	IdUser         uint
	ProductName    string
	ProductQty     int
	Price          int
	ProductPicture string
}

func FromDomain(dp domain.Core) Cart {
	return Cart{
		Model:          gorm.Model{ID: dp.ID},
		IdProduct:      dp.IdProduct,
		IdUser:         dp.IdUser,
		ProductName:    dp.ProductName,
		ProductQty:     dp.ProductQty,
		Price:          dp.Price,
		ProductPicture: dp.ProductPicture,
	}
}

func ToDomain(dp Cart) domain.Core {
	return domain.Core{
		ID:             dp.ID,
		IdProduct:      dp.IdProduct,
		IdUser:         dp.IdUser,
		ProductName:    dp.ProductName,
		ProductQty:     dp.ProductQty,
		Price:          dp.Price,
		ProductPicture: dp.ProductPicture,
	}
}

func ToDomainArray(dp []Cart) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{ID: val.ID, IdProduct: val.IdProduct, IdUser: val.IdUser, ProductName: val.ProductName, ProductQty: val.ProductQty, Price: val.Price, ProductPicture: val.ProductPicture})
	}
	return res
}
