// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	executor "doggos/executor"

	mock "github.com/stretchr/testify/mock"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

// Execute provides a mock function with given fields: req
func (_m *Executor) Execute(req *executor.Request) (*executor.Response, error) {
	ret := _m.Called(req)

	var r0 *executor.Response
	if rf, ok := ret.Get(0).(func(*executor.Request) *executor.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*executor.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}