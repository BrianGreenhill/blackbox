// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

// Schedule provides a mock function with given fields: schedule
func (_m *Scheduler) Schedule(schedule string) error {
	ret := _m.Called(schedule)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(schedule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
