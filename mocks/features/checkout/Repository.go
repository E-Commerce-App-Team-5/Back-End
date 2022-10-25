// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "ecommerce/features/checkout/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id uint) (domain.Core, error) {
	ret := _m.Called(id)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Repository) Get(id uint) ([]domain.Core, error) {
	ret := _m.Called(id)

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func(uint) []domain.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
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

// Insert provides a mock function with given fields: newHistory, newCheckout
func (_m *Repository) Insert(newHistory []domain.HistoryCore, newCheckout domain.Core) (domain.Core, error) {
	ret := _m.Called(newHistory, newCheckout)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func([]domain.HistoryCore, domain.Core) domain.Core); ok {
		r0 = rf(newHistory, newCheckout)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]domain.HistoryCore, domain.Core) error); ok {
		r1 = rf(newHistory, newCheckout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newCheckout
func (_m *Repository) Update(newCheckout domain.Core) error {
	ret := _m.Called(newCheckout)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Core) error); ok {
		r0 = rf(newCheckout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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