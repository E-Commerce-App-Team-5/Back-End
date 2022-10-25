package delivery

import (
	"ecommerce/features/history_order/domain"
)

func SuccessResponse(msg string, data interface{}) interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type BuyResponse struct {
	ID         uint `json:"id"`
	IdCheckout uint `json:"id_checkout"`
	IdProduct  uint `json:"id_product"`
	ProductQty int  `json:"product_qty"`
	Price      int  `json:"price"`
	ProductName string `json:"product_name"`
	NamaToko string `json:"nama_toko"`
}

func ToResponse(input interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "buy":
		var arr []BuyResponse
		cnv := input.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, BuyResponse{ID: val.ID, IdCheckout: val.IdCheckout, IdProduct: val.IdProduct, ProductName: val.ProductName, NamaToko: val.NamaToko, ProductQty: val.ProductQty, Price: val.Price})
		}
		res = arr
	}
	return res
}
