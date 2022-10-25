package services

import (
	"ecommerce/features/checkout/domain"
	mocks "ecommerce/mocks/features/checkout"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCheckout(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add Checkout", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.Anything).Return(domain.Core{ID: uint(1), IdPembeli: 1, OrderId: "Order-101", Token: "213", Link: "http.youtube.com"}, nil).Once()

		srv := New(repo)
		coreInput := domain.Core{IdPembeli: uint(1), OrderId: "Order-101", Token: "213", Link: "http.youtube.com"}
		input := []domain.HistoryCore{{IdPembeli: uint(1), IdProduct: 3, Price: 20000}}
		res, err := srv.AddCheckout(input, coreInput)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add Checkout", func(t *testing.T) {
		repo.On("Insert", mock.Anything, mock.Anything).Return(domain.Core{}, errors.New("error add checkout")).Once()
		srv := New(repo)
		res, err := srv.AddCheckout([]domain.HistoryCore{}, domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteCheckout(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete Checkout", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), IdPembeli: 1, OrderId: "Order-101", Token: "213", Link: "http.youtube.com"}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteCheckout(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete Checkout", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.DeleteCheckout(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetCheckout(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get Checkout", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{{ID: uint(1), IdPembeli: 1, OrderId: "Order-101", Token: "213", Link: "http.youtube.com"}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetCheckout(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Checkout", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetCheckout(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
