package delivery

import (
	"ecommerce/features/user/domain"
)

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
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	NamaToko    string `json:"nama_toko"`
	UserPicture string `json:"user_picture"`
}

type RegisterResponse struct {
	Username string `json:"usename"`
	Email    string `json:"email"`
}

type GetResponse struct {
	ID          uint             `json:"id"`
	Username    string           `json:"username"`
	Email       string           `json:"email"`
	Fullname    string           `json:"fullname"`
	Phone       string           `json:"phone"`
	NamaToko    string           `json:"nama_toko"`
	UserPicture string           `json:"user_picture"`
	Product     []domain.Product `json:"products"`
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	UserPicture string `json:"user_picture"`
	Token       string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Fullname: cnv.Fullname, UserPicture: cnv.UserPicture, Token: cnv.Token}
	case "register":
		cnv := core.(domain.Core)
		res = RegisterResponse{Username: cnv.Username, Email: cnv.Email}
	case "update":
		cnv := core.(domain.Core)
		res = UpdateResponse{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Fullname: cnv.Fullname, UserPicture: cnv.UserPicture, Phone: cnv.Phone, NamaToko: cnv.NamaToko}
	}
	return res
}

func ToResponseGetUser(user interface{}, product interface{}) interface{} {
	var res interface{}
	var resProduct []domain.Product
	cnvProduct := product.([]domain.Product)

	for _, val := range cnvProduct {
		resProduct = append(resProduct, domain.Product{ID: val.ID, IdUser: val.IdUser, NamaToko: val.NamaToko,
			ProductName: val.ProductName, ProductDetail: val.ProductDetail, ProductQty: val.ProductQty, Price: val.Price, ProductPicture: val.ProductPicture})
	}

	cnvUser := user.(domain.Core)
	resUser := GetResponse{ID: cnvUser.ID, Username: cnvUser.Username, Email: cnvUser.Email, Fullname: cnvUser.Fullname, Phone: cnvUser.Phone,
		NamaToko: cnvUser.NamaToko, UserPicture: cnvUser.UserPicture, Product: resProduct}

	res = resUser
	return res
}
