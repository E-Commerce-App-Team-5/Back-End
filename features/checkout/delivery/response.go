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

type GetResponse struct {
	ID             uint   `json:"id"`
	IdProduct      uint   `json:"id_product"`
	IdUser         uint   `json:"id_user"`
	NamaToko       string `json:"nama_toko"`
	ProductName    string `json:"product_name"`
	ProductQty     int    `json:"product_qty"`
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
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
type UpdateResponse struct {
	ID         uint `json:"id"`
	ProductQty int  `json:"product_qty"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{ID: cnv.ID, OrderId: cnv.OrderId, GrossAmount: cnv.GrossAmount, Token: cnv.Token, Link: cnv.Link}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID}
	}

	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetResponse{ID: cnv.ID})
	}
	res = arr
	return res
}
