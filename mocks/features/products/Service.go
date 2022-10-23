// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "ecommerce/features/products/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: newProduct
func (_m *Service) AddProduct(newProduct domain.Core) (domain.Core, error) {
	ret := _m.Called(newProduct)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newProduct)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProduct provides a mock function with given fields: ID
func (_m *Service) DeleteProduct(ID uint) (domain.Core, error) {
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

// GetProduct provides a mock function with given fields: page
func (_m *Service) GetProduct(page int) ([]domain.Core, error) {
	ret := _m.Called(page)

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func(int) []domain.Core); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: input
func (_m *Service) UpdateProduct(input domain.Core) (domain.Core, error) {
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

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
