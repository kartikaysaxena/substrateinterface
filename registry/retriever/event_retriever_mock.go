// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package retriever

import (
	parser "github.com/kartikaysaxena/susbtrateinterface/registry/parser"
	types "github.com/kartikaysaxena/susbtrateinterface/types"
	mock "github.com/stretchr/testify/mock"
)

// EventRetrieverMock is an autogenerated mock type for the EventRetriever type
type EventRetrieverMock struct {
	mock.Mock
}

// GetEvents provides a mock function with given fields: blockHash
func (_m *EventRetrieverMock) GetEvents(blockHash types.Hash) ([]*parser.Event, error) {
	ret := _m.Called(blockHash)

	var r0 []*parser.Event
	if rf, ok := ret.Get(0).(func(types.Hash) []*parser.Event); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*parser.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEventRetrieverMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventRetrieverMock creates a new instance of EventRetrieverMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventRetrieverMock(t NewEventRetrieverMockT) *EventRetrieverMock {
	mock := &EventRetrieverMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
