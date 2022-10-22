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
	ID             uint
	IdUser         uint
	NamaToko       string
	ProductName    string
	ProductDetail  string
	ProductQty     int
	Price          int
	ProductPicture string
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
