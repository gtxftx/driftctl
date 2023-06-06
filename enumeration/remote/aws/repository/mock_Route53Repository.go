// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	route53 "github.com/aws/aws-sdk-go/service/route53"
	mock "github.com/stretchr/testify/mock"
)

// MockRoute53Repository is an autogenerated mock type for the Route53Repository type
type MockRoute53Repository struct {
	mock.Mock
}

// ListAllHealthChecks provides a mock function with given fields:
func (_m *MockRoute53Repository) ListAllHealthChecks() ([]*route53.HealthCheck, error) {
	ret := _m.Called()

	var r0 []*route53.HealthCheck
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*route53.HealthCheck, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*route53.HealthCheck); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*route53.HealthCheck)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllZones provides a mock function with given fields:
func (_m *MockRoute53Repository) ListAllZones() ([]*route53.HostedZone, error) {
	ret := _m.Called()

	var r0 []*route53.HostedZone
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*route53.HostedZone, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*route53.HostedZone); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*route53.HostedZone)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecordsForZone provides a mock function with given fields: zoneId
func (_m *MockRoute53Repository) ListRecordsForZone(zoneId string) ([]*route53.ResourceRecordSet, error) {
	ret := _m.Called(zoneId)

	var r0 []*route53.ResourceRecordSet
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*route53.ResourceRecordSet, error)); ok {
		return rf(zoneId)
	}
	if rf, ok := ret.Get(0).(func(string) []*route53.ResourceRecordSet); ok {
		r0 = rf(zoneId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*route53.ResourceRecordSet)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(zoneId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRoute53Repository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRoute53Repository creates a new instance of MockRoute53Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRoute53Repository(t mockConstructorTestingTNewMockRoute53Repository) *MockRoute53Repository {
	mock := &MockRoute53Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
