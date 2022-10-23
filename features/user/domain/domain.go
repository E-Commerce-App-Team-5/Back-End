package domain

type Core struct {
	ID          uint
	Fullname    string
	Username    string
	Email       string
	Password    string
	UserPicture string
	Phone       string
	DOB         string
	NamaToko    string
	Token       string
	ProductDetail ProductDetail
}

type Product struct {
	ID             uint `json:"id"`
	IdUser         uint	`json:"id_user"`
	NamaToko       string `json:"nama_toko"`
	ProductName    string `json:"product_name"`
	ProductDetail  string `json:"product_detail"`
	ProductQty     int `json:"product_qty"`
	Price          int `json:"price"`
	ProductPicture string `json:"profile_picture"`
}

type ProductDetail struct {
	ProductDetail []Product
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newUser Core) (Core, error)
	Get(username string) (Core, error)
	Delete(ID uint) (Core, error)
	Edit(input Core) (Core, error)
	Login(input Core) (Core, error)
	GetProduct(id uint) ([]Product, error)
}

type Service interface { // Bisnis logic
	Register(newUser Core) (Core, error)
	GetUser(username string) (Core, []Product, error)
	DeleteUser(ID uint) (Core, error)
	UpdateUser(input Core) (Core, error)
	Login(input Core) (Core, string, error)
}
