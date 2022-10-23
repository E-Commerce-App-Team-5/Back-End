package services

import (
	"errors"
	"ecommerce/features/user/domain"
	mocks "ecommerce/mocks/features/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Password: "$2a$10$szpOHiZl0Uvv.Wr1hTAwKeSbTb2E2igPNzPHqW.C0u5xMmLRomaYS "}, nil).Once()
		srv := New(repo)
		input := domain.Core{Username: "fatur", Password: "fatur123"}
		res, _, err := srv.Login(input)
		assert.NotEmpty(t, res)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Wrong username Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Password: "asgfasg"}, errors.New("username not found")).Once()
		srv := New(repo)
		input := domain.Core{Username: "fatur", Password: "fatur123"}
		res, _, err := srv.Login(input)
		assert.Empty(t, res)
		assert.EqualError(t, err, "username not found")
		repo.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Username: "fatur22", Email: "fatur@gmail.com", 
		Password: "fatur123"}, nil).Once()

		srv := New(repo)
		input := domain.Core{Password: "fatur123", Username: "faturfawkes", Email: "fatur@gmail.com"}
		res, err := srv.Register(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add User", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.Register(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), Username: "fatur00", Email: "fatur28@yahoo.com",Password: "fatur123"}, nil).Once()
		srv := New(repo)
		res, err := srv.DeleteUser(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete User", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.DeleteUser(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{ID: uint(1), Username: "fatur", Fullname: "fatur fawkes"}, nil).Once()
		product := domain.Product{NamaToko: "juragan99" ,ProductName: "monitor"}
		repo.On("GetProduct", mock.Anything).Return([]domain.Product{product}, nil).Once()
		srv := New(repo)
		resUser, resProduct, err := srv.GetUser("fatur")
		assert.Nil(t, err)
		assert.NotEmpty(t, resUser)
		assert.NotEmpty(t, resProduct)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{}, errors.New("no data")).Once()
		// repo.On("GetProduct", mock.Anything).Return(nil, errors.New("get product error")).Once()
		srv := New(repo)
		resUser, resProduct, err := srv.GetUser("fatur")
		assert.NotNil(t, err)
		assert.Empty(t, resUser)
		assert.Empty(t, resProduct)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Get User", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(domain.Core{ID: uint(1), Username: "fatur", Fullname: "fatur fawkes"}, nil).Once()
		repo.On("GetProduct", mock.Anything).Return([]domain.Product{}, errors.New("get product error")).Once()
		srv := New(repo)
		resUser, resProduct, err := srv.GetUser("fatur")
		assert.NotNil(t, err)
		assert.Empty(t, resUser)
		assert.Empty(t, resProduct)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything).Return(domain.Core{ID: 1, Username: "username", Password: "fatur123",
			Email: "fatur@gmail.com"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: 1, Username: "username", Fullname: "fatur rohman", Password: "fatur123", Phone: "-8900001",
			Email: "fatur@gmail.com"}
		res, err := srv.UpdateUser(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update User", func(t *testing.T) {
		repo.On("Edit", mock.Anything).Return(domain.Core{}, errors.New("error update user")).Once()
		srv := New(repo)
		var input domain.Core
		res, err := srv.UpdateUser(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
