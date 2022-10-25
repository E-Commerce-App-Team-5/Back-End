package repository

import (
	"ecommerce/features/history_order/domain"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	IdCheckout  uint
	IdProduct   uint
	ProductQty  int
	Price       int
	// IdPembeli   uint   `gorm:"-:migration"`
	ProductName string `gorm:"-:migration"`
	NamaToko  string `gorm:"-:migration"`
}

func ToDomainArray(cnv []History) []domain.Core {
	var res []domain.Core
	for _, val := range cnv {
		res = append(res, domain.Core{
			ID:         val.ID,
			IdCheckout: val.IdCheckout,
			IdProduct:  val.IdProduct,
			ProductName: val.ProductName,
			NamaToko: val.NamaToko,
			ProductQty: val.ProductQty,
			Price:      val.Price,
		})
	}
	return res
}
