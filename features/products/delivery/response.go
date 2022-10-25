package delivery

import "ecommerce/features/products/domain"

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

type UpdateResponse struct {
	ID             uint   `json:"id"`
	ProductName    string `json:"product_name"`
	ProductDetail  string `json:"product_detail"`
	ProductQty     int    `json:"product_qty"`
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
}

type RegisterResponse struct {
	IdUser         uint   `json:"id_user"`
	ProductName    string `json:"product_name"`
	ProductQty     int    `json:"product_qty"`
	Price          int    `json:"price"`
	ProductPicture string `json:"product_picture"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{IdUser: cnv.IdUser, ProductName: cnv.ProductName, 
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, ProductName: cnv.ProductName, ProductDetail: cnv.ProductDetail,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture}
	}
	return res
}

func ToResponseProduct(core interface{}, code string) interface{} {
	var res interface{}
	var arr []UpdateResponse
	val := core.([]domain.Core)
	for _, cnv := range val {
		arr = append(arr, UpdateResponse{ID: cnv.ID, ProductName: cnv.ProductName, ProductDetail: cnv.ProductDetail,
			ProductQty: cnv.ProductQty, Price: cnv.Price, ProductPicture: cnv.ProductPicture})
	}
	res = arr
	return res
}
