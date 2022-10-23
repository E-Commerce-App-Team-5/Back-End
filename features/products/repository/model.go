package repository

import (
	"ecommerce/features/products/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductDetail  string
	ProductQty     int
	Price          int
	ProductPicture string
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
