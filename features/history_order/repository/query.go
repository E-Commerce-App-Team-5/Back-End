package repository

import (
	"ecommerce/features/history_order/domain"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repostory {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) GetBuy(id uint) ([]domain.Core, error) {
	var res []History
	if err := rq.db.Select("histories.id", "histories.id_product", "id_checkout", "histories.product_qty" ,"price").Joins("left join checkouts on checkouts.id = histories.id_checkout").Where("checkouts.id_pembeli = ?", id).Find(&res).Scan(&res).Error; err != nil {
		return []domain.Core{}, err
	}
	cnv := ToDomainArray(res)
	return cnv, nil
}

func (rq *repoQuery) GetSell(id uint) ([]domain.Core, error) {
	return []domain.Core{}, nil
}