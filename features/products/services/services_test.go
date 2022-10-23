package services

import (
	"ecommerce/features/products/domain"
	mocks "ecommerce/mocks/features/products"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add Product", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}, nil).Once()

		srv := New(repo)
		input := domain.Core{IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10, Price: 10000, ProductPicture: "srv.jpg"}
		res, err := srv.AddProduct(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add Products", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.AddProduct(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteProduct(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.DeleteProduct(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{{ID: uint(1), IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10,
			Price: 10000, ProductPicture: "srv.jpg"}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetProduct(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetProduct(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything).Return(domain.Core{ID: 1, IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10, Price: 10000, ProductPicture: "srv.jpg"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: 1, IdUser: uint(1), NamaToko: "tokosebek", ProductName: "buah", ProductQty: 10, Price: 10000, ProductPicture: "srv.jpg"}
		res, err := srv.UpdateProduct(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything).Return(domain.Core{}, errors.New("error update user")).Once()
		srv := New(repo)
		var input domain.Core
		res, err := srv.UpdateProduct(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
