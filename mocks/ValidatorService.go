// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	validators "github.com/manicar2093/charly_team_api/validators"
	mock "github.com/stretchr/testify/mock"
)

// ValidatorService is an autogenerated mock type for the ValidatorService type
type ValidatorService struct {
	mock.Mock
}

// Validate provides a mock function with given fields: e
func (_m *ValidatorService) Validate(e interface{}) validators.ValidateOutput {
	ret := _m.Called(e)

	var r0 validators.ValidateOutput
	if rf, ok := ret.Get(0).(func(interface{}) validators.ValidateOutput); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(validators.ValidateOutput)
	}

	return r0
}
