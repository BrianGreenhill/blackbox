// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	internal "greenhill/backend/pkg/internal"

	mock "github.com/stretchr/testify/mock"
)

// Notifier is an autogenerated mock type for the Notifier type
type Notifier struct {
	mock.Mock
}

// Notify provides a mock function with given fields: _a0
func (_m *Notifier) Notify(_a0 *internal.CheckResult) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.CheckResult) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
