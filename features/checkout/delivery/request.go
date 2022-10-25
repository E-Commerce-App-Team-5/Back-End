package delivery

import "ecommerce/features/checkout/domain"

type RegisterFormat struct {
	IdPembeli  uint    `json:"id_pembeli" form:"id_pembeli"`
	IdProduct  uint    `json:"id_product" form:"id_product"`
	ProductQty int     `json:"product_qty" form:"product_qty"`
	Price      float32 `json:"price" form:"price"`
}

type CheckoutFormat struct {
	IdPembeli uint `json:"id_pembeli" form:"id_pembeli"`
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

func ToDomainHistory(i interface{}) []domain.HistoryCore {
	var arr []domain.HistoryCore
	val := i.([]RegisterFormat)
	for _, cnv := range val {
		arr = append(arr, domain.HistoryCore{IdPembeli: cnv.IdPembeli, IdProduct: cnv.IdProduct, ProductQty: cnv.ProductQty, Price: int(cnv.Price)})
	}
	return arr
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case CheckoutFormat:
		cnv := i.(CheckoutFormat)
		return domain.Core{IdPembeli: cnv.IdPembeli}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID}
	}
	return domain.Core{}
}
