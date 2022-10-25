package repository

import (
	"ecommerce/features/checkout/domain"

	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	IdPembeli   uint
	OrderId     string `gorm:"unique"`
	GrossAmount float32
	Token       string
	Link        string
	Status      string
	Historys    []History `gorm:"foreignKey:IdCheckout"`
}

type History struct {
	gorm.Model
	IdCheckout uint
	IdProduct  uint
	ProductQty int
	Price      int
	Status     string
}

func FromDomainHistory(dp []domain.HistoryCore) []History {
	var res []History
	for _, val := range dp {
		res = append(res, History{Model: gorm.Model{ID: val.ID},
			IdProduct:  val.IdProduct,
			ProductQty: val.ProductQty,
			Price:      val.Price,
			Status:     val.Status})
	}
	return res
}

func FromDomain(dp domain.Core) Checkout {
	return Checkout{
		Model:       gorm.Model{ID: dp.ID},
		IdPembeli:   dp.IdPembeli,
		OrderId:     dp.OrderId,
		GrossAmount: dp.GrossAmount,
		Token:       dp.Token,
		Link:        dp.Link,
		Status:      dp.Status,
	}
}

func ToDomain(dp Checkout) domain.Core {
	return domain.Core{
		ID:          dp.ID,
		IdPembeli:   dp.IdPembeli,
		OrderId:     dp.OrderId,
		GrossAmount: dp.GrossAmount,
		Token:       dp.Token,
		Link:        dp.Link,
		Status:      dp.Status,
	}
}

func ToDomainArray(dp []Checkout) []domain.Core {
	var res []domain.Core
	for _, val := range dp {
		res = append(res, domain.Core{ID: val.ID, IdPembeli: val.IdPembeli,
			OrderId:     val.OrderId,
			GrossAmount: val.GrossAmount,
			Token:       val.Token,
			Link:        val.Link,
			Status:      val.Status})
	}
	return res
}
