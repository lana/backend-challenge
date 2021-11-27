// Code generated by mockery v2.9.4. DO NOT EDIT.

package storagemocks

import (
	context "context"
	models "patriciabonaldy/lana/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, basketID, productCode
func (_m *Repository) AddProduct(ctx context.Context, basketID string, productCode string) (models.Basket, error) {
	ret := _m.Called(ctx, basketID, productCode)

	var r0 models.Basket
	if rf, ok := ret.Get(0).(func(context.Context, string, string) models.Basket); ok {
		r0 = rf(ctx, basketID, productCode)
	} else {
		r0 = ret.Get(0).(models.Basket)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, basketID, productCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBasket provides a mock function with given fields: ctx, id
func (_m *Repository) CreateBasket(ctx context.Context, id string) (models.Basket, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Basket
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Basket); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Basket)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBasketByID provides a mock function with given fields: ctx, id
func (_m *Repository) FindBasketByID(ctx context.Context, id string) (models.Basket, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Basket
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Basket); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Basket)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveProduct provides a mock function with given fields: ctx, basketID, productCode
func (_m *Repository) RemoveProduct(ctx context.Context, basketID string, productCode string) (models.Basket, error) {
	ret := _m.Called(ctx, basketID, productCode)

	var r0 models.Basket
	if rf, ok := ret.Get(0).(func(context.Context, string, string) models.Basket); ok {
		r0 = rf(ctx, basketID, productCode)
	} else {
		r0 = ret.Get(0).(models.Basket)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, basketID, productCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
