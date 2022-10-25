package repository

import (
	"ecommerce/features/products/domain"

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
	NamaToko       string `gorm:"-:migration" gorm:"<-"`
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
	NamaToko       string `gorm:"-:migration" gorm:"->"`
	ProductName    string `gorm:"-:migration" gorm:"->"`
	ProductQty     int
	Price          int    `gorm:"-:migration" gorm:"->"`
	ProductPicture string `gorm:"-:migration" gorm:"->"`
}

func FromDomain(dp domain.Core) Product {
	return Product{
		Model:          gorm.Model{ID: dp.ID},
		IdUser:         dp.IdUser,
		NamaToko:       dp.NamaToko,
		ProductName:    dp.ProductName,
		ProductDetail:  dp.ProductDetail,
		ProductQty:     dp.ProductQty,
		Price:          dp.Price,
		ProductPicture: dp.ProductPicture,
	}
}

func ToDomain(dp Product) domain.Core {
	return domain.Core{
		ID:             dp.ID,
		IdUser:         dp.IdUser,
		NamaToko:       dp.NamaToko,
		ProductName:    dp.ProductName,
		ProductDetail:  dp.ProductDetail,
		ProductQty:     dp.ProductQty,
		Price:          dp.Price,
		ProductPicture: dp.ProductPicture,
	}
}

func ToDomainArray(dp []Product) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{ID: val.ID, IdUser: val.IdUser, NamaToko: val.NamaToko, ProductName: val.ProductName,
			ProductDetail: val.ProductDetail, ProductQty: val.ProductQty, Price: val.Price, ProductPicture: val.ProductPicture})
	}
	return res
}
