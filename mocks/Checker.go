// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	internal "github.com/briangreenhill/blackbox/internal"
	mock "github.com/stretchr/testify/mock"
)

// Checker is an autogenerated mock type for the Checker type
type Checker struct {
	mock.Mock
}

// DoCheck provides a mock function with given fields: _a0
func (_m *Checker) DoCheck(_a0 internal.Target) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.Target) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
