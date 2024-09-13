// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	state "github.com/kartikaysaxena/substrateinterface/rpc/state"
	types "github.com/kartikaysaxena/substrateinterface/types"
	mock "github.com/stretchr/testify/mock"
)

// State is an autogenerated mock type for the State type
type State struct {
	mock.Mock
}

// GetChildKeys provides a mock function with given fields: childStorageKey, prefix, blockHash
func (_m *State) GetChildKeys(childStorageKey types.StorageKey, prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error) {
	ret := _m.Called(childStorageKey, prefix, blockHash)

	var r0 []types.StorageKey
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, types.Hash) []types.StorageKey); ok {
		r0 = rf(childStorageKey, prefix, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(childStorageKey, prefix, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildKeysLatest provides a mock function with given fields: childStorageKey, prefix
func (_m *State) GetChildKeysLatest(childStorageKey types.StorageKey, prefix types.StorageKey) ([]types.StorageKey, error) {
	ret := _m.Called(childStorageKey, prefix)

	var r0 []types.StorageKey
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey) []types.StorageKey); ok {
		r0 = rf(childStorageKey, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(childStorageKey, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorage provides a mock function with given fields: childStorageKey, key, target, blockHash
func (_m *State) GetChildStorage(childStorageKey types.StorageKey, key types.StorageKey, target interface{}, blockHash types.Hash) (bool, error) {
	ret := _m.Called(childStorageKey, key, target, blockHash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, interface{}, types.Hash) bool); ok {
		r0 = rf(childStorageKey, key, target, blockHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, interface{}, types.Hash) error); ok {
		r1 = rf(childStorageKey, key, target, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageHash provides a mock function with given fields: childStorageKey, key, blockHash
func (_m *State) GetChildStorageHash(childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (types.Hash, error) {
	ret := _m.Called(childStorageKey, key, blockHash)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, types.Hash) types.Hash); ok {
		r0 = rf(childStorageKey, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageHashLatest provides a mock function with given fields: childStorageKey, key
func (_m *State) GetChildStorageHashLatest(childStorageKey types.StorageKey, key types.StorageKey) (types.Hash, error) {
	ret := _m.Called(childStorageKey, key)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey) types.Hash); ok {
		r0 = rf(childStorageKey, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageLatest provides a mock function with given fields: childStorageKey, key, target
func (_m *State) GetChildStorageLatest(childStorageKey types.StorageKey, key types.StorageKey, target interface{}) (bool, error) {
	ret := _m.Called(childStorageKey, key, target)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, interface{}) bool); ok {
		r0 = rf(childStorageKey, key, target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, interface{}) error); ok {
		r1 = rf(childStorageKey, key, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageRaw provides a mock function with given fields: childStorageKey, key, blockHash
func (_m *State) GetChildStorageRaw(childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(childStorageKey, key, blockHash)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(childStorageKey, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageRawLatest provides a mock function with given fields: childStorageKey, key
func (_m *State) GetChildStorageRawLatest(childStorageKey types.StorageKey, key types.StorageKey) (*types.StorageDataRaw, error) {
	ret := _m.Called(childStorageKey, key)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey) *types.StorageDataRaw); ok {
		r0 = rf(childStorageKey, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageSize provides a mock function with given fields: childStorageKey, key, blockHash
func (_m *State) GetChildStorageSize(childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (types.U64, error) {
	ret := _m.Called(childStorageKey, key, blockHash)

	var r0 types.U64
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey, types.Hash) types.U64); ok {
		r0 = rf(childStorageKey, key, blockHash)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageSizeLatest provides a mock function with given fields: childStorageKey, key
func (_m *State) GetChildStorageSizeLatest(childStorageKey types.StorageKey, key types.StorageKey) (types.U64, error) {
	ret := _m.Called(childStorageKey, key)

	var r0 types.U64
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.StorageKey) types.U64); ok {
		r0 = rf(childStorageKey, key)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeys provides a mock function with given fields: prefix, blockHash
func (_m *State) GetKeys(prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error) {
	ret := _m.Called(prefix, blockHash)

	var r0 []types.StorageKey
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.Hash) []types.StorageKey); ok {
		r0 = rf(prefix, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.Hash) error); ok {
		r1 = rf(prefix, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeysLatest provides a mock function with given fields: prefix
func (_m *State) GetKeysLatest(prefix types.StorageKey) ([]types.StorageKey, error) {
	ret := _m.Called(prefix)

	var r0 []types.StorageKey
	if rf, ok := ret.Get(0).(func(types.StorageKey) []types.StorageKey); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: blockHash
func (_m *State) GetMetadata(blockHash types.Hash) (*types.Metadata, error) {
	ret := _m.Called(blockHash)

	var r0 *types.Metadata
	if rf, ok := ret.Get(0).(func(types.Hash) *types.Metadata); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Metadata)
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

// GetMetadataLatest provides a mock function with given fields:
func (_m *State) GetMetadataLatest() (*types.Metadata, error) {
	ret := _m.Called()

	var r0 *types.Metadata
	if rf, ok := ret.Get(0).(func() *types.Metadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Metadata)
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

// GetRuntimeVersion provides a mock function with given fields: blockHash
func (_m *State) GetRuntimeVersion(blockHash types.Hash) (*types.RuntimeVersion, error) {
	ret := _m.Called(blockHash)

	var r0 *types.RuntimeVersion
	if rf, ok := ret.Get(0).(func(types.Hash) *types.RuntimeVersion); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RuntimeVersion)
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

// GetRuntimeVersionLatest provides a mock function with given fields:
func (_m *State) GetRuntimeVersionLatest() (*types.RuntimeVersion, error) {
	ret := _m.Called()

	var r0 *types.RuntimeVersion
	if rf, ok := ret.Get(0).(func() *types.RuntimeVersion); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RuntimeVersion)
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

// GetStorage provides a mock function with given fields: key, target, blockHash
func (_m *State) GetStorage(key types.StorageKey, target interface{}, blockHash types.Hash) (bool, error) {
	ret := _m.Called(key, target, blockHash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.StorageKey, interface{}, types.Hash) bool); ok {
		r0 = rf(key, target, blockHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, interface{}, types.Hash) error); ok {
		r1 = rf(key, target, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageHash provides a mock function with given fields: key, blockHash
func (_m *State) GetStorageHash(key types.StorageKey, blockHash types.Hash) (types.Hash, error) {
	ret := _m.Called(key, blockHash)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.Hash) types.Hash); ok {
		r0 = rf(key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.Hash) error); ok {
		r1 = rf(key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageHashLatest provides a mock function with given fields: key
func (_m *State) GetStorageHashLatest(key types.StorageKey) (types.Hash, error) {
	ret := _m.Called(key)

	var r0 types.Hash
	if rf, ok := ret.Get(0).(func(types.StorageKey) types.Hash); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageLatest provides a mock function with given fields: key, target
func (_m *State) GetStorageLatest(key types.StorageKey, target interface{}) (bool, error) {
	ret := _m.Called(key, target)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.StorageKey, interface{}) bool); ok {
		r0 = rf(key, target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, interface{}) error); ok {
		r1 = rf(key, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageRaw provides a mock function with given fields: key, blockHash
func (_m *State) GetStorageRaw(key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(key, blockHash)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.Hash) error); ok {
		r1 = rf(key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageRawLatest provides a mock function with given fields: key
func (_m *State) GetStorageRawLatest(key types.StorageKey) (*types.StorageDataRaw, error) {
	ret := _m.Called(key)

	var r0 *types.StorageDataRaw
	if rf, ok := ret.Get(0).(func(types.StorageKey) *types.StorageDataRaw); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageSize provides a mock function with given fields: key, blockHash
func (_m *State) GetStorageSize(key types.StorageKey, blockHash types.Hash) (types.U64, error) {
	ret := _m.Called(key, blockHash)

	var r0 types.U64
	if rf, ok := ret.Get(0).(func(types.StorageKey, types.Hash) types.U64); ok {
		r0 = rf(key, blockHash)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey, types.Hash) error); ok {
		r1 = rf(key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageSizeLatest provides a mock function with given fields: key
func (_m *State) GetStorageSizeLatest(key types.StorageKey) (types.U64, error) {
	ret := _m.Called(key)

	var r0 types.U64
	if rf, ok := ret.Get(0).(func(types.StorageKey) types.U64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.StorageKey) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorage provides a mock function with given fields: keys, startBlock, block
func (_m *State) QueryStorage(keys []types.StorageKey, startBlock types.Hash, block types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(keys, startBlock, block)

	var r0 []types.StorageChangeSet
	if rf, ok := ret.Get(0).(func([]types.StorageKey, types.Hash, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(keys, startBlock, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.StorageKey, types.Hash, types.Hash) error); ok {
		r1 = rf(keys, startBlock, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageAt provides a mock function with given fields: keys, block
func (_m *State) QueryStorageAt(keys []types.StorageKey, block types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(keys, block)

	var r0 []types.StorageChangeSet
	if rf, ok := ret.Get(0).(func([]types.StorageKey, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(keys, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.StorageKey, types.Hash) error); ok {
		r1 = rf(keys, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageAtLatest provides a mock function with given fields: keys
func (_m *State) QueryStorageAtLatest(keys []types.StorageKey) ([]types.StorageChangeSet, error) {
	ret := _m.Called(keys)

	var r0 []types.StorageChangeSet
	if rf, ok := ret.Get(0).(func([]types.StorageKey) []types.StorageChangeSet); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.StorageKey) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageLatest provides a mock function with given fields: keys, startBlock
func (_m *State) QueryStorageLatest(keys []types.StorageKey, startBlock types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(keys, startBlock)

	var r0 []types.StorageChangeSet
	if rf, ok := ret.Get(0).(func([]types.StorageKey, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(keys, startBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.StorageKey, types.Hash) error); ok {
		r1 = rf(keys, startBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeRuntimeVersion provides a mock function with given fields:
func (_m *State) SubscribeRuntimeVersion() (*state.RuntimeVersionSubscription, error) {
	ret := _m.Called()

	var r0 *state.RuntimeVersionSubscription
	if rf, ok := ret.Get(0).(func() *state.RuntimeVersionSubscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.RuntimeVersionSubscription)
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

// SubscribeStorageRaw provides a mock function with given fields: keys
func (_m *State) SubscribeStorageRaw(keys []types.StorageKey) (*state.StorageSubscription, error) {
	ret := _m.Called(keys)

	var r0 *state.StorageSubscription
	if rf, ok := ret.Get(0).(func([]types.StorageKey) *state.StorageSubscription); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.StorageSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.StorageKey) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewStateT interface {
	mock.TestingT
	Cleanup(func())
}

// NewState creates a new instance of State. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewState(t NewStateT) *State {
	mock := &State{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
