package services

import (
	"ecommerce/features/user/domain"
	"ecommerce/middlewares"
	"errors"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

func (us *userService) Login(input domain.Core) (domain.Core, string, error) {
	res, err := us.qry.Login(input)
	if err != nil {
		log.Error(err.Error(), "email not found")
		return domain.Core{}, "", errors.New("email not found")
	}

	pass := domain.Core{Password: res.Password}
	check := bcrypt.CompareHashAndPassword([]byte(pass.Password), []byte(input.Password))
	if check != nil { 
		log.Error(check, " wrong password")
		return domain.Core{}, "", errors.New("wrong password")
	}
	token, err := middlewares.CreateToken(int(res.ID))

	return res, token, err
}

func (us *userService) UpdateUser(input domain.Core) (domain.Core, error) {
	orgPass := input.Password
	if input.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error(err.Error())
			return domain.Core{}, errors.New("cannot encrypt password")
		}
		input.Password = string(generate)
	}
	res, err := us.qry.Edit(input)
	if err != nil {
		return domain.Core{}, err
	}
	res.Password = orgPass
	return res, nil
}

func (us *userService) DeleteUser(id uint) (domain.Core, error) {
	res, err := us.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}

func (us *userService) Register(newUser domain.Core) (domain.Core, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error(err.Error())
		return domain.Core{}, errors.New("cannot encrypt password")
	}
	orgPass := newUser.Password
	newUser.Password = string(generate)
	res, err := us.qry.Insert(newUser)
	if err != nil {
		return domain.Core{}, errors.New("user has registered")
	}

	res.Password = orgPass

	return res, nil
}

func (us *userService) GetUser(username string) (domain.Core, []domain.Product, error) {
	resUser, err := us.qry.Get(username)
	if err != nil {
		// log.Error(err.Error())
		return domain.Core{}, []domain.Product{}, errors.New("no data")
	}

	resProduct, err := us.qry.GetProduct(resUser.ID)
	if err != nil {
		return domain.Core{}, []domain.Product{}, errors.New("get product error")
	}
	return resUser, resProduct, nil
}
