// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "schedule/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// GetById provides a mock function with given fields: ctx, userId, scheduleId
func (_m *Repo) GetById(ctx context.Context, userId int64, scheduleId int) (*entity.Schedule, error) {
	ret := _m.Called(ctx, userId, scheduleId)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *entity.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) (*entity.Schedule, error)); ok {
		return rf(ctx, userId, scheduleId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int) *entity.Schedule); ok {
		r0 = rf(ctx, userId, scheduleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int) error); ok {
		r1 = rf(ctx, userId, scheduleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUser provides a mock function with given fields: ctx, userId
func (_m *Repo) GetByUser(ctx context.Context, userId int64) ([]entity.Schedule, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetByUser")
	}

	var r0 []entity.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.Schedule, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.Schedule); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *Repo) Save(ctx context.Context, _a1 *entity.Schedule) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Schedule) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
