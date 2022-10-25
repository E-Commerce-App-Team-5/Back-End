package services

import (
	"ecommerce/features/checkout/domain"
	"errors"

	"github.com/labstack/gommon/log"
)

type checkoutService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &checkoutService{
		qry: repo,
	}
}

func (cs *checkoutService) DeleteCheckout(id uint) (domain.Core, error) {
	res, err := cs.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}
	return res, err
}

func (cs *checkoutService) AddCheckout(newHistory []domain.HistoryCore, newChekout domain.Core) (domain.Core, error) {
	res, err := cs.qry.Insert(newHistory, newChekout)
	if err != nil {
		return domain.Core{}, errors.New("some problem on database")
	}

	return res, nil
}

func (cs *checkoutService) GetCheckout(id uint) ([]domain.Core, error) {
	res, err := cs.qry.Get(id)
	if err != nil {
		log.Error(err.Error())
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}

func (cs *checkoutService) UpdateCheckout(newCheckout domain.Core) {
	err := cs.qry.Update(newCheckout)
	if err != nil {
		log.Error(err.Error())
	}
}
