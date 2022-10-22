package repository

import (
	"ecommerce/features/user/domain"
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

func (rq *repoQuery) AddPhotos(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Model(&cnv).Where("id = ?", input.ID).Update("images", cnv.Fullname).Error; err != nil {
		log.Fatal("error update data")
		return domain.Core{}, err
	}
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Login(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Where("email = ?", cnv.Email).First(&cnv).Error; err != nil {
		log.Fatal("error get data")
		return domain.Core{}, err
	}
	log.Print(cnv.ID, "ini id")
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	if err := rq.db.Where("id = ?", id).Delete(&User{}); err != nil {
		return domain.Core{}, err.Error
	}
	return domain.Core{}, nil
}

func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) Edit(input domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(input)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	input = ToDomain(cnv)
	return input, nil
}

func (rq *repoQuery) Get(username string) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "username = ?", username).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) GetProduct(id uint) ([]domain.Product, error) {
	var resProduct []Product

	if err := rq.db.Where("id_user = ?", id).Find(&resProduct).Error; err != nil {
		return []domain.Product{}, err
	}

	Product := ToDomainArray(resProduct)
	return Product, nil
}
