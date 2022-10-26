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
	ID             uint ``
	IdProduct      uint
	ProductQty     int
	NamaToko       string
	ProductName    string
	PriceSum       int
	ProductPicture string
}

func ToResponse(code string, history []domain.Core) interface{} {
	var res interface{}

	switch code {
	case "buy":
		var arr []BuyResponse
		for _, val := range  history {
			arr = append(arr, BuyResponse{
				ID: val.ID,
				IdProduct: val.IdProduct,
				ProductQty: val.ProductQty,
				NamaToko: val.NamaToko,
				ProductName: val.ProductName,
				PriceSum: val.PriceSum,
				ProductPicture: val.ProductPicture,
			})
		}
		res = arr
	}
	return res
}
