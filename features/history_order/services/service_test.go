package services

import (
	"ecommerce/features/history_order/domain"
	mocks "ecommerce/mocks/features/history_order/domain"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHistoryBuy(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get", func(t *testing.T) {
		repo.On("GetBuy", mock.Anything).Return([]domain.Core{domain.Core{
			ID:             1,
			IdProduct:      2,
			NamaPembeli:    "ujang",
			NamaToko:       "toko_store",
			ProductName:    "kursi gaming",
			ProductQty:     3,
			ProductsDetail: "ini kursi gaming keren",
			Price:          600000,
			ProductPicture: "www.photo.com",
		}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetBuy(1)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get", func(t *testing.T) {
		repo.On("GetBuy", mock.Anything).Return([]domain.Core{}, errors.New("error get data")).Once()
		srv := New(repo)
		res, err := srv.GetBuy(1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestHistorySell(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get", func(t *testing.T) {
		repo.On("GetSell", mock.Anything).Return([]domain.Core{domain.Core{
			ID:             1,
			IdProduct:      2,
			NamaPembeli:    "ujang",
			NamaToko:       "toko_store",
			ProductName:    "kursi gaming",
			ProductQty:     3,
			ProductsDetail: "ini kursi gaming keren",
			Price:          600000,
			ProductPicture: "www.photo.com",
		}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetSell(1)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get", func(t *testing.T) {
		repo.On("GetSell", mock.Anything).Return([]domain.Core{}, errors.New("error get data")).Once()
		srv := New(repo)
		res, err := srv.GetSell(1)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
