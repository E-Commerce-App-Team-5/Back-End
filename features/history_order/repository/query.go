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

	if err := rq.db.Model(&[]History{}).Select("histories.id", "products.product_name", "histories.product_qty", 
	"histories.price", "histories.id_product","products.product_name", "products.product_detail","users.nama_toko", "products.product_picture").
	Joins("left join products on products.id = histories.id_product").
	Joins("left join checkouts on checkouts.id = histories.id_checkout").
	Joins("left join users on users.id = products.id_user").
	Where("checkouts.id_pembeli = ? AND checkouts.status = ?", int(id), "settlement").
	Scan(&res).Find(&res).
	Error; err != nil {
		log.Print("error query")
		return []domain.Core{}, err
	}


	cnv := ToDomain(res)
	return cnv, nil
}

func (rq *repoQuery) GetSell(id uint) ([]domain.Core, error) { 
	var res []History
	if err := rq.db.Model(&[]History{}).Select("histories.id, histories.id_product", "histories.product_qty", "histories.price",
	"users.username", "products.product_name", "products.product_detail", "products.product_picture").
	Joins("left join checkouts on checkouts.id = histories.id_checkout").
	Joins("left join products on products.id = histories.id_product").
	Joins("left join users on users.id = checkouts.id_pembeli").
	Where("products.id_user = ? AND checkouts.status = ?", id, "settlement").
	Find(&res).Scan(&res).Error; err != nil {
		log.Print("error query")
		return []domain.Core{}, err
	}
	cnv := ToDomain(res)
	return cnv, nil
}