// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package tokenclaimsgenerator

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTokenClaimsGenerator is an autogenerated mock type for the TokenClaimsGenerator type
type MockTokenClaimsGenerator struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx, req
func (_m *MockTokenClaimsGenerator) Run(ctx context.Context, req *TokenClaimsGeneratorRequest) (*TokenClaimsGeneratorResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *TokenClaimsGeneratorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *TokenClaimsGeneratorRequest) *TokenClaimsGeneratorResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TokenClaimsGeneratorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *TokenClaimsGeneratorRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
