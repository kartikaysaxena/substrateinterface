// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	author "github.com/kartikaysaxena/substrateinterface/rpc/author"
	extrinsic "github.com/kartikaysaxena/substrateinterface/types/extrinsic"

	mock "github.com/stretchr/testify/mock"

	types "github.com/kartikaysaxena/substrateinterface/types"
)

// Author is an autogenerated mock type for the Author type
type Author struct {
	mock.Mock
}

// PendingExtrinsics provides a mock function with given fields:
func (_m *Author) PendingExtrinsics() ([]extrinsic.Extrinsic, error) {
	ret := _m.Called()

	var r0 []extrinsic.Extrinsic
	if rf, ok := ret.Get(0).(func() []extrinsic.Extrinsic); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]extrinsic.Extrinsic)
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

// SubmitAndWatchExtrinsic provides a mock function with given fields: xt
func (_m *Author) SubmitAndWatchExtrinsic(xt extrinsic.Extrinsic) (*author.ExtrinsicStatusSubscription, error) {
	ret := _m.Called(xt)

	var r0 *author.ExtrinsicStatusSubscription
	if rf, ok := ret.Get(0).(func(extrinsic.Extrinsic) *author.ExtrinsicStatusSubscription); ok {
		r0 = rf(xt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*author.ExtrinsicStatusSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(extrinsic.Extrinsic) error); ok {
		r1 = rf(xt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitExtrinsic provides a mock function with given fields: xt
func (_m *Author) SubmitExtrinsic(xt extrinsic.Extrinsic) (types.Hash, error) {
	ret := _m.Called(xt)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(extrinsic.Extrinsic) types.Hash); ok {
		r0 = rf(xt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(extrinsic.Extrinsic) error); ok {
		r1 = rf(xt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewAuthorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthor creates a new instance of Author. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthor(t NewAuthorT) *Author {
	mock := &Author{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
