// Code generated by mockery v2.10.0. DO NOT EDIT.

package repository

import (
	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	mock "github.com/stretchr/testify/mock"
)

// MockELBV2Repository is an autogenerated mock type for the ELBV2Repository type
type MockELBV2Repository struct {
	mock.Mock
}

// ListAllLoadBalancers provides a mock function with given fields:
func (_m *MockELBV2Repository) ListAllLoadBalancers() ([]*elbv2.LoadBalancer, error) {
	ret := _m.Called()

	var r0 []*elbv2.LoadBalancer
	if rf, ok := ret.Get(0).(func() []*elbv2.LoadBalancer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*elbv2.LoadBalancer)
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

// ListLoadBalancerListeners provides a mock function with given fields: _a0
func (_m *MockELBV2Repository) ListAllLoadBalancerListeners(_a0 string) ([]*elbv2.Listener, error) {
	ret := _m.Called(_a0)

	var r0 []*elbv2.Listener
	if rf, ok := ret.Get(0).(func(string) []*elbv2.Listener); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*elbv2.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
