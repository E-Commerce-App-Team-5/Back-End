package delivery

import (
	"ecommerce/features/products/domain"
)

type RegisterFormat struct {
	IdUser         uint   `json:"id_user" form:"id_user"`
	ProductName    string `json:"product_name" form:"product_name"`
	ProductDetail  string `json:"product_detail" form:"product_detail"`
	ProductQty     int    `json:"product_qty" form:"product_qty"`
	Price          int    `json:"price" form:"price"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type UpdateFormat struct {
	ID             uint   `json:"id" form:"id"`
	IdUser         uint   `json:"id_user" form:"id_user"`
	ProductName    string `json:"product_name" form:"product_name"`
	ProductDetail  string `json:"product_detail" form:"product_detail"`
	ProductQty     int    `json:"product_qty" form:"product_qty"`
	Price          int    `json:"price" form:"price"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{IdUser: cnv.IdUser, ProductName: cnv.ProductName, ProductDetail: cnv.ProductDetail,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, IdUser: cnv.IdUser, ProductName: cnv.ProductName, ProductDetail: cnv.ProductDetail,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	}
	return domain.Core{}
}
