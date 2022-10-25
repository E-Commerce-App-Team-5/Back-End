package repository

import (
	"ecommerce/features/history_order/domain"
	"log"

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
	if err := rq.db.Select("histories.id", "histories.id_product", "histories.id_checkout", "histories.product_qty" ,"histories.price", "products.product_name", "users.nama_toko").
	Joins("left join checkouts on checkouts.id = histories.id_checkout").
	Joins("left join products on products.id = histories.id_product").
	Joins("left join users on users.id = products.id_user").
	Where("checkouts.id_pembeli = ?", id).
	Find(&res).Scan(&res).Error; err != nil {
		return []domain.Core{}, err
	}

	log.Print(res[0].ProductName)
	log.Print(res[0].NamaToko)
	cnv := ToDomainArray(res)
	return cnv, nil
}

func (rq *repoQuery) GetSell(id uint) ([]domain.Core, error) {
	return []domain.Core{}, nil
}