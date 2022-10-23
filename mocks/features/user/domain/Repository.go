// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "ecommerce/features/user/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ID
func (_m *Repository) Delete(ID uint) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Edit provides a mock function with given fields: input
func (_m *Repository) Edit(input domain.Core) (domain.Core, error) {
	ret := _m.Called(input)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: username
func (_m *Repository) Get(username string) (domain.Core, error) {
	ret := _m.Called(username)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(string) domain.Core); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: id
func (_m *Repository) GetProduct(id uint) ([]domain.Product, error) {
	ret := _m.Called(id)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(uint) []domain.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newUser
func (_m *Repository) Insert(newUser domain.Core) (domain.Core, error) {
	ret := _m.Called(newUser)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: input
func (_m *Repository) Login(input domain.Core) (domain.Core, error) {
	ret := _m.Called(input)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
