package delivery

import "ecommerce/features/cart/domain"

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
	ProductDetail  string `json:"product_detail"`
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
}

type RegisterResponse struct {
	IdProduct  uint `json:"id_product"`
	ProductQty int  `json:"product_qty"`
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
		res = RegisterResponse{IdProduct: cnv.IdProduct, ProductQty: cnv.ProductQty}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, ProductQty: cnv.ProductQty}
	}

	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetResponse{ID: cnv.ID, IdProduct: cnv.IdProduct, ProductDetail: cnv.ProductDetail,IdUser: cnv.IdUser, NamaToko: cnv.NamaToko, ProductName: cnv.ProductName,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture})
	}
	res = arr
	return res
}
