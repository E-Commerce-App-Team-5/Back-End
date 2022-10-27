package services

import (
	"ecommerce/features/history_order/domain"
	"errors"
)

type productService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &productService{
		qry: repo,
	}
}

func (ps *productService) GetBuy(id uint) ([]domain.Core, error) {
	res, err := ps.qry.GetBuy(id)
	if err != nil {
		return []domain.Core{}, errors.New("There is problem on server.")
	}
	for i, val := range res {
		res[i].PriceSum = val.Price * val.ProductQty
	}
	return res,  nil
}

func (ps *productService) GetSell(id uint) ([]domain.Core, error){
	res, err := ps.qry.GetSell(id)
	if err != nil {
		return []domain.Core{}, errors.New("There is problem on server.")
	}
	for i, val := range res {
		res[i].PriceSum = val.Price
	}
	return res, nil
}