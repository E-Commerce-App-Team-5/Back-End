package services

import (
	"ecommerce/features/cart/domain"
	mocks "ecommerce/mocks/features/cart"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCart(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add Cart", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), IdProduct: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}, nil).Once()

		srv := New(repo)
		input := domain.Core{IdProduct: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10, Price: 10000, ProductPicture: "srv.jpg"}
		res, err := srv.AddCart(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add Cart", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.AddCart(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete Cart", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteCart(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete Cart", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.DeleteCart(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get Cart", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{{ID: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetCart(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Cart", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetCart(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}
