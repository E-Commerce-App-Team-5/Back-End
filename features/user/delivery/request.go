package delivery

import (
	"ecommerce/features/user/domain"
)

type RegisterFormat struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UpdateFormat struct {
	ID           uint
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
	Fullname     string `form:"fullname" json:"fullname"`
	Phone        string `form:"phone" json:"phone"`
	NamaToko    string `form:"nama_toko" json:"nama_toko"`
	UserPicture string `form:"user_picture" json:"user_picture"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type GetId struct {
	id uint `param:"id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password}
	case GetId:
		cnv := i.(GetId)
		return domain.Core{ID: cnv.id}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Username: cnv.Username, Email: cnv.Email, Password: cnv.Password, Fullname: cnv.Fullname,
			Phone: cnv.Phone, NamaToko: cnv.NamaToko, UserPicture: cnv.UserPicture}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.Core{Email: cnv.Email, Password: cnv.Password}
	}
	return domain.Core{}
}
