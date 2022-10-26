package repository

import (
	"ecommerce/features/history_order/domain"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	IdProduct      uint
	ProductQty     int
	Price          int
	ProductName    string `gorm:"-:migration" gorm:"<-"`
	NamaToko       string `gorm:"-:migration" gorm:"<-"`
	ProductPicture string `gorm:"-:migration" gorm:"<-"`
}

func ToDomain(cnv []History) []domain.Core {
	var res []domain.Core
	for _, val := range cnv {
		res = append(res, domain.Core{
			ID:          val.ID,
			IdProduct:   val.IdProduct,
			ProductQty:  val.ProductQty,
			Price:       val.Price,
			NamaToko:    val.NamaToko,
			ProductName: val.ProductName,
			ProductPicture: val.ProductPicture,
		})
	}
	return res
}
