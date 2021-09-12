// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Filterable is an autogenerated mock type for the Filterable type
type Filterable struct {
	mock.Mock
}

// GetFilter provides a mock function with given fields: filterName
func (_m *Filterable) GetFilter(filterName string) error {
	ret := _m.Called(filterName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filterName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *Filterable) Run() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetContext provides a mock function with given fields: ctx
func (_m *Filterable) SetContext(ctx context.Context) {
	_m.Called(ctx)
}

// SetValues provides a mock function with given fields: values
func (_m *Filterable) SetValues(values interface{}) {
	_m.Called(values)
}
