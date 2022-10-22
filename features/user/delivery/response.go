package delivery

import "ecommerce/features/user/domain"

func SuccessResponseWithData(msg string, data interface{}) map[string]interface{} {
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
	ID       uint   `json:"id"`
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email"`
	Fullname     string `form:"fullname" json:"fullname"`
	Phone        string `form:"phone" json:"phone"`
	NamaToko    string `form:"nama_toko" json:"nama_toko"`
	UserPicture string `form:"user_picture" json:"user_picture"`
	domain.ProductDetail
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	UserPicture string `json:""`
	Token       string `json:"token"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "get":
		cnv := core.(domain.Core)
		res = GetResponse{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Fullname: cnv.Fullname, UserPicture: cnv.UserPicture}
	case "login":
		cnv := core.(domain.Core)
		res = LoginResponse{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Fullname: cnv.Fullname, UserPicture: cnv.UserPicture, Token: cnv.Token}
	}
	return res
}

func ToResponseGetUser(user interface{}, product interface{}, code string) interface{} {
	var res interface{}
	var resProduct []domain.Product
	cnvProduct := product.([]domain.Product)
	for _, val := range cnvProduct {
		resProduct = append(resProduct, domain.Product{ID: val.ID, IdUser: val.IdUser, NamaToko: val.NamaToko,
		ProductName: val.ProductName, ProductDetail: val.ProductDetail, ProductQty: val.ProductQty, Price: val.Price, ProductPicture: val.ProductPicture})	
	}

	cnvUser := user.(domain.Core)
	res = GetResponse{ID: cnvUser.ID, Username: cnvUser.Username, Email: cnvUser.Email, Fullname: cnvUser.Fullname, Phone: cnvUser.Phone,
		NamaToko: cnvUser.NamaToko, UserPicture: cnvUser.UserPicture, ProductDetail: domain.ProductDetail{resProduct}}

	return res
}
