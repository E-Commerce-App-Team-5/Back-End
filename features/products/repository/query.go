package repository

import (
	"ecommerce/features/products/domain"
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
	if err := rq.db.Where("id = ?", id).Delete(&Product{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newProduct domain.Core) (domain.Core, error) {
	var cnv Product = FromDomain(newProduct)
	if err := rq.db.Select("id_user", "product_name", "product_detail", "product_qty", "product_picture", "price").Create(&cnv).Error; err != nil {
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
		if err := rq.db.Model(&Product{}).Limit(20).Select("products.id, products.product_name, products.product_detail, products.product_qty, products.price, products.product_picture ,users.nama_toko").Order("products.created_at desc").Joins("left join users on users.id = products.id_user").Scan(&resQry).Error; err != nil {
			return nil, err
		}
	} else {
		i := page * 20
		if err := rq.db.Model(&Product{}).Offset(i).Limit(20).Select("products.id, products.product_name, products.product_detail, products.product_qty, products.price, products.product_picture ,users.nama_toko").Order("products.created_at desc").Joins("left join users on users.id = products.id_user").Scan(&resQry).Error; err != nil {
			return nil, err
		}
	}

	log.Print("ini dari log ", resQry)
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
