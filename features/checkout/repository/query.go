package repository

import (
	"ecommerce/features/checkout/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	if err := rq.db.Where("id = ?", id).Delete(&Checkout{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newHistory []domain.HistoryCore, newCheckout domain.Core) (domain.Core, error) {
	var res Checkout = FromDomain(newCheckout)
	var cnv []History = FromDomainHistory(newHistory)

	if err := rq.db.Create(&res).Error; err != nil {
		return domain.Core{}, err
	}

	for i := 0; i < len(cnv); i++ {
		cnv[i].IdCheckout = res.ID
		if err := rq.db.Create(&cnv[i]).Error; err != nil {
			return domain.Core{}, err
		}
	}

	// selesai dari DB
	newCheckout = ToDomain(res)
	return newCheckout, nil
}

func (rq *repoQuery) Get(id uint) ([]domain.Core, error) {
	var resQry []Checkout
	if err := rq.db.Where("id_pembeli=? AND status='pending'", id).Find(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
