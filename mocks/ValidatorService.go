// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ValidatorService is an autogenerated mock type for the ValidatorService type
type ValidatorService struct {
	mock.Mock
}

// Validate provides a mock function with given fields: e
func (_m *ValidatorService) Validate(e interface{}) (bool, error) {
	ret := _m.Called(e)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
