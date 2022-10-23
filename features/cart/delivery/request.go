package delivery

import "ecommerce/features/cart/domain"

type RegisterFormat struct {
	IdProduct      uint   `json:"id_product" form:"id_product"`
	IdUser         uint   `json:"id_user" form:"id_user"`
	NamaToko       string `json:"nama_toko" form:"nama_toko"`
	ProductName    string `json:"product_name" form:"product_name"`
	ProductQty     int    `json:"product_qty" form:"product_qty"`
	Price          int    `json:"price" form:"price"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
}

type UpdateFormat struct {
	ID             uint   `json:"id" form:"id"`
	IdProduct      uint   `json:"id_product" form:"id_product"`
	IdUser         uint   `json:"id_user" form:"id_user"`
	NamaToko       string `json:"nama_toko" form:"nama_toko"`
	ProductName    string `json:"product_name" form:"product_name"`
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
		return domain.Core{IdProduct: cnv.IdProduct, IdUser: cnv.IdUser, NamaToko: cnv.NamaToko, ProductName: cnv.ProductName,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, IdProduct: cnv.IdProduct, IdUser: cnv.IdUser, NamaToko: cnv.NamaToko, ProductName: cnv.ProductName,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	}
	return domain.Core{}
}
