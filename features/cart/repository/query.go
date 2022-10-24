package repository

import (
	"ecommerce/features/cart/domain"

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
	if err := rq.db.Where("id = ?", id).Delete(&Cart{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newCart domain.Core) (domain.Core, error) {
	var cnv Cart = FromDomain(newCart)

	if err := rq.db.Select("id_product", "id_user", "product_qty").Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	newCart = ToDomain(cnv)
	return newCart, nil
}

func (rq *repoQuery) Edit(input domain.Core) (domain.Core, error) {
	var cnv Cart = FromDomain(input)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Get(id uint) ([]domain.Core, error) {
	var resQry []Cart
	if err := rq.db.Where("carts.id_user=?", id).Find(&resQry).Joins("left join products on products.id = carts.id_product").Joins("left join users on users.id = carts.id_user").Scan(&resQry).Error; err != nil {
		return nil, err
	}

	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
