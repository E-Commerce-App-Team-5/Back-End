package services

import "ecommerce/features/history_order/domain"

type productService struct {
	qry domain.Repostory
}

func New(repo domain.Repostory) domain.Services {
	return &productService{
		qry: repo,
	}
}

func (ps *productService) GetBuy(id uint) ([]domain.Core, error) {
	res, err := ps.qry.GetBuy(id)
	if err != nil {
		return []domain.Core{}, err
	}
	return res, nil
}

func (ps *productService) GetSell(id uint) ([]domain.Core, error){
	return []domain.Core{}, nil
}