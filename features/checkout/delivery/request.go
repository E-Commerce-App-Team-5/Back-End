package delivery

import "ecommerce/features/checkout/domain"

type RegisterFormat struct {
	IdPembeli  uint    `json:"id_pembeli" form:"id_pembeli"`
	IdProduct  uint    `json:"id_product" form:"id_product"`
	ProductQty int     `json:"product_qty" form:"product_qty"`
	PriceSum   float32 `json:"price_sum" form:"price_sum"`
}

type CheckoutFormat struct {
	IdPembeli uint    `json:"id_pembeli" form:"id_pembeli"`
	Price     float32 `json:"price" form:"price"`
}

func ToDomainHistory(i interface{}, j interface{}) ([]domain.HistoryCore, domain.Core) {
	var arr []domain.HistoryCore
	res := j.(CheckoutFormat)
	val := i.([]RegisterFormat)
	for _, cnv := range val {
		arr = append(arr, domain.HistoryCore{IdPembeli: cnv.IdPembeli, IdProduct: cnv.IdProduct, ProductQty: cnv.ProductQty, Price: int(cnv.PriceSum)})
		res.Price += cnv.PriceSum
	}
	resCheckout := domain.Core{IdPembeli: res.IdPembeli, GrossAmount: res.Price}
	return arr, resCheckout
}
