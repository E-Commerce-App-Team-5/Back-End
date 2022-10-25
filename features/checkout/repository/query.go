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
	var res Checkout
	var cnv []History
	for i := 0; i < len(newHistory); i++ {
		rq.db.Where("id_product=?", newHistory[i].IdProduct).Select("price").First(&newHistory[i])
		newHistory[i].Price = newHistory[i].Price * newHistory[i].ProductQty
		res.GrossAmount += float32(newHistory[i].Price)
	}
	cnv = FromDomainHistory(newHistory)
	for i := 0; i < len(cnv); i++ {
		if err := rq.db.Create(&cnv[i]).Error; err != nil {
			return domain.Core{}, err
		}
	}

	if err := rq.db.Create(&res).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	newCheckout = ToDomain(res)
	return newCheckout, nil
}

func (rq *repoQuery) Edit(input domain.Core) (domain.Core, error) {
	var cnv Checkout = FromDomain(input)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Get(id uint) ([]domain.Core, error) {
	var resQry []Checkout
	if err := rq.db.Where("carts.id_user=?", id).Find(&resQry).Joins("left join products on products.id = carts.id_product").Joins("left join users on users.id = carts.id_user").Scan(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
