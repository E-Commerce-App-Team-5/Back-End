package services

import (
	"ecommerce/features/cart/domain"
	"errors"

	"github.com/labstack/gommon/log"
)

type productService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &productService{
		qry: repo,
	}
}

func (ps *productService) DeleteCart(id uint) (domain.Core, error) {
	res, err := ps.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}

func (ps *productService) AddCart(newProduct domain.Core) (domain.Core, error) {
	res, err := ps.qry.Insert(newProduct)
	if err != nil {
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (ps *productService) GetCart(id uint) ([]domain.Core, error) {
	res, err := ps.qry.Get(id)
	if err != nil {
		log.Error(err.Error())
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}
