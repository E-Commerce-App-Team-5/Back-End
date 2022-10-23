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
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
}

type RegisterResponse struct {
	IdProduct      uint   `json:"id_product"`
	IdUser         uint   `json:"id_user"`
	NamaToko       string `json:"nama_toko"`
	ProductName    string `json:"product_name"`
	ProductQty     int    `json:"product_qty"`
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	cnv := core.(domain.Core)
	res = RegisterResponse{IdProduct: cnv.IdProduct, IdUser: cnv.IdUser, NamaToko: cnv.NamaToko, ProductName: cnv.ProductName,
		ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []GetResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, GetResponse{ID: cnv.ID, IdProduct: cnv.IdProduct, IdUser: cnv.IdUser, NamaToko: cnv.NamaToko, ProductName: cnv.ProductName,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture})
	}
	res = arr
	return res
}