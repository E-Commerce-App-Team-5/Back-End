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
	ID             uint
	IdProduct      uint   `json:"id_product"`
	ProductQty     int    `json:"product_qty"`
	NamaToko       string `json:"nama_toko"`
	ProductName    string `json:"prodcut_name"`
	PriceSum       int    `json:"price_sum"`
	ProductPicture string `json:"product_picture"`
}

type SellResponse struct {
	ID             uint
	IdProduct      uint   `json:"id_product"`
	ProductQty     int    `json:"product_qty"`
	NamaPembeli    string `json:"nama_pembeli"`
	ProductName    string `json:"prodcut_name"`
	PriceSum       int    `json:"price_sum"`
	ProductPicture string `json:"product_picture"`
}

func ToResponse(code string, history []domain.Core) interface{} {
	var res interface{}

	switch code {
	case "buy":
		var arr []BuyResponse
		for _, val := range history {
			arr = append(arr, BuyResponse{
				ID:             val.ID,
				IdProduct:      val.IdProduct,
				ProductQty:     val.ProductQty,
				NamaToko:       val.NamaToko,
				ProductName:    val.ProductName,
				PriceSum:       val.PriceSum,
				ProductPicture: val.ProductPicture,
			})
		}
		res = arr
	case "sell":
		var arr []SellResponse
		for _, val := range history {
			arr = append(arr, SellResponse{
				ID:             val.ID,
				IdProduct:      val.IdProduct,
				ProductQty:     val.ProductQty,
				ProductName:    val.ProductName,
				NamaPembeli:    val.NamaPembeli,
				PriceSum:       val.PriceSum,
				ProductPicture: val.ProductPicture,
			})
		}
		res = arr
	}
	return res
}
