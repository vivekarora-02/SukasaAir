// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "sukasaair/dto"

	mock "github.com/stretchr/testify/mock"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *Database) Client() *mongo.Client {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 *mongo.Client
	if rf, ok := ret.Get(0).(func() *mongo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Client)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: name
func (_m *Database) Collection(name string) *mongo.Collection {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Collection")
	}

	var r0 *mongo.Collection
	if rf, ok := ret.Get(0).(func(string) *mongo.Collection); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.Collection)
		}
	}

	return r0
}

// DeleteMany provides a mock function with given fields: ctx, filter
func (_m *Database) DeleteMany(ctx context.Context, filter interface{}) error {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOne provides a mock function with given fields: ctx, filter
func (_m *Database) FindOne(ctx context.Context, filter interface{}) (*dto.Seat, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 *dto.Seat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*dto.Seat, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *dto.Seat); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Seat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitSeats provides a mock function with given fields:
func (_m *Database) InitSeats() {
	_m.Called()
}

// InsertMany provides a mock function with given fields: ctx, documents
func (_m *Database) InsertMany(ctx context.Context, documents []interface{}) error {
	ret := _m.Called(ctx, documents)

	if len(ret) == 0 {
		panic("no return value specified for InsertMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) error); ok {
		r0 = rf(ctx, documents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertOne provides a mock function with given fields: ctx, document
func (_m *Database) InsertOne(ctx context.Context, document interface{}) error {
	ret := _m.Called(ctx, document)

	if len(ret) == 0 {
		panic("no return value specified for InsertOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, document)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMany provides a mock function with given fields: ctx, filter, update
func (_m *Database) UpdateMany(ctx context.Context, filter interface{}, update interface{}) error {
	ret := _m.Called(ctx, filter, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, filter, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOne provides a mock function with given fields: ctx, filter, update
func (_m *Database) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	ret := _m.Called(ctx, filter, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}) error); ok {
		r0 = rf(ctx, filter, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
