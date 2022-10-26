package repository

import (
	"ecommerce/features/checkout/domain"
	"log"

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

	log.Print()
	if err := rq.db.Create(&res).Error; err != nil {
		return domain.Core{}, err
	}

	for i := 0; i < len(cnv); i++ {
		cnv[i].IdCheckout = res.ID
		if err := rq.db.Create(&cnv[i]).Error; err != nil {
			return domain.Core{}, err
		}
		if err := rq.db.Where("id_product=? AND id_user=? AND product_qty=?", cnv[i].IdProduct, res.IdPembeli, cnv[i].ProductQty).Delete(&Cart{}).Error; err != nil {
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

func (rq *repoQuery) Update(newCheckout domain.Core) error {
	var cnv Checkout = FromDomain(newCheckout)
	if err := rq.db.Where("order_id=?", newCheckout.OrderId).Updates(&cnv).Error; err != nil {
		return err
	}

	var res []History
	var produk, temp Product
	rq.db.Where("order_id=?", newCheckout.OrderId).First(&cnv)
	rq.db.Where("id_checkout=?", &cnv.ID).Find(&res)
	for _, val := range res {
		rq.db.Where("id=?", val.IdProduct).First(&produk)
		temp.ProductQty = produk.ProductQty - val.ProductQty
		rq.db.Where("id=?", val.IdProduct).Updates(&temp)
	}
	newCheckout = ToDomain(cnv)
	return nil
}
