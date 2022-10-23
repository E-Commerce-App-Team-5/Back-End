package repository

import (
	"ecommerce/features/products/domain"

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
	if err := rq.db.Where("id = ?", id).Delete(&Product{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newProduct domain.Core) (domain.Core, error) {
	var cnv Product = FromDomain(newProduct)
	var tempUser User
	rq.db.Where("id=?", cnv.IdUser).First(&tempUser)
	cnv.NamaToko = tempUser.NamaToko
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	newProduct = ToDomain(cnv)
	return newProduct, nil
}

func (rq *repoQuery) Edit(input domain.Core) (domain.Core, error) {
	var cnv Product = FromDomain(input)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Get(page int) ([]domain.Core, error) {
	var resQry []Product
	if page == 0 {
		if err := rq.db.Limit(20).Order("created_at desc").Find(&resQry).Error; err != nil {
			return nil, err
		}
	} else {
		i := page * 20
		if err := rq.db.Offset(i).Limit(20).Order("created_at desc").Find(&resQry).Error; err != nil {
			return nil, err
		}
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
