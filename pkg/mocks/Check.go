// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	internal "github.com/briangreenhill/blackbox/pkg/internal"

	mock "github.com/stretchr/testify/mock"
)

// Check is an autogenerated mock type for the Check type
type Check struct {
	mock.Mock
}

// DoCheck provides a mock function with given fields: _a0, _a1
func (_m *Check) DoCheck(_a0 internal.Target, _a1 internal.Notifier) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.Target, internal.Notifier) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}