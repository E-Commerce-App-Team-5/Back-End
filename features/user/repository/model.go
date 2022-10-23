package repository

import (
	"ecommerce/features/user/domain"

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
	NamaToko       string
	ProductName    string
	ProductQty     int
	Price          int
	ProductPicture string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:       gorm.Model{ID: du.ID},
		Fullname:    du.Fullname,
		Username:    du.Username,
		Email:       du.Email,
		Password:    du.Password,
		UserPicture: du.UserPicture,
		DOB:         du.DOB,
		Phone:       du.Phone,
		NamaToko:    du.NamaToko,
	}
}

func ToDomain(du User) domain.Core {
	return domain.Core{
		ID:          du.ID,
		Fullname:    du.Fullname,
		Username:    du.Username,
		Email:       du.Email,
		Password:    du.Password,
		UserPicture: du.UserPicture,
		DOB:         du.DOB,
		Phone:       du.Phone,
		NamaToko:    du.NamaToko,
	}
}

func ToDomainArray(dp []Product) []domain.Product {
	var res []domain.Product
	for _, val := range dp {
		res = append(res, domain.Product{ID: val.ID, IdUser: val.IdUser, NamaToko: val.NamaToko, ProductName: val.ProductName,
			ProductDetail: val.ProductDetail, ProductQty: val.ProductQty, Price: val.Price, ProductPicture: val.ProductPicture})
	}
	return res
}
