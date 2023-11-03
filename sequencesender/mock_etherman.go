// Code generated by mockery v2.28.1. DO NOT EDIT.

package sequencesender

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "github.com/ethereum/go-ethereum/core/types"

	etherman "github.com/0xPolygonHermez/zkevm-node/etherman"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygonHermez/zkevm-node/etherman/types"
)

// ethermanMock is an autogenerated mock type for the ethermanInterface type
type ethermanMock struct {
	mock.Mock
}

// BuildSequenceBatchesTxData provides a mock function with given fields: sender, sequences, l2Coinbase, committeeSignaturesAndAddrs
func (_m *ethermanMock) BuildSequenceBatchesTxData(sender common.Address, sequences []types.Sequence, l2Coinbase common.Address, committeeSignaturesAndAddrs []byte) (*common.Address, []byte, error) {
	ret := _m.Called(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)

	var r0 *common.Address
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address, []byte) (*common.Address, []byte, error)); ok {
		return rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address, []byte) *common.Address); ok {
		r0 = rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address, []byte) []byte); ok {
		r1 = rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(common.Address, []types.Sequence, common.Address, []byte) error); ok {
		r2 = rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EstimateGasSequenceBatches provides a mock function with given fields: sender, sequences, l2Coinbase, committeeSignaturesAndAddrs
func (_m *ethermanMock) EstimateGasSequenceBatches(sender common.Address, sequences []types.Sequence, l2Coinbase common.Address, committeeSignaturesAndAddrs []byte) (*coretypes.Transaction, error) {
	ret := _m.Called(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)

	var r0 *coretypes.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address, []byte) (*coretypes.Transaction, error)); ok {
		return rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	}
	if rf, ok := ret.Get(0).(func(common.Address, []types.Sequence, common.Address, []byte) *coretypes.Transaction); ok {
		r0 = rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, []types.Sequence, common.Address, []byte) error); ok {
		r1 = rf(sender, sequences, l2Coinbase, committeeSignaturesAndAddrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentDataCommittee provides a mock function with given fields:
func (_m *ethermanMock) GetCurrentDataCommittee() (*etherman.DataCommittee, error) {
	ret := _m.Called()

	var r0 *etherman.DataCommittee
	var r1 error
	if rf, ok := ret.Get(0).(func() (*etherman.DataCommittee, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *etherman.DataCommittee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*etherman.DataCommittee)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchTimestamp provides a mock function with given fields:
func (_m *ethermanMock) GetLastBatchTimestamp() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBatchNumber provides a mock function with given fields:
func (_m *ethermanMock) GetLatestBatchNumber() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlockTimestamp provides a mock function with given fields: ctx
func (_m *ethermanMock) GetLatestBlockTimestamp(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewEthermanMock interface {
	mock.TestingT
	Cleanup(func())
}

// newEthermanMock creates a new instance of ethermanMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newEthermanMock(t mockConstructorTestingTnewEthermanMock) *ethermanMock {
	mock := &ethermanMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
