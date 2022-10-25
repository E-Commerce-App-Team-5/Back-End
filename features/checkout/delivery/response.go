package delivery

import "ecommerce/features/checkout/domain"

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

type RegisterResponse struct {
	ID          uint    `json:"id"`
	IdPembeli   uint    `json:"id_pembeli"`
	OrderId     string  `json:"order_id"`
	GrossAmount float32 `json:"gross_amount"`
	Token       string  `json:"token"`
	Link        string  `json:"link"`
	Status      string  `json:"status"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, IdPembeli: cnv.IdPembeli, OrderId: cnv.OrderId, GrossAmount: cnv.GrossAmount, Token: cnv.Token, Link: cnv.Link}
	}

	return res
}
