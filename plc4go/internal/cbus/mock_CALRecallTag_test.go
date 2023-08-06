/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.3. DO NOT EDIT.

package cbus

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	readwritemodel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockCALRecallTag is an autogenerated mock type for the CALRecallTag type
type MockCALRecallTag struct {
	mock.Mock
}

type MockCALRecallTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCALRecallTag) EXPECT() *MockCALRecallTag_Expecter {
	return &MockCALRecallTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockCALRecallTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCALRecallTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockCALRecallTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetAddressString() *MockCALRecallTag_GetAddressString_Call {
	return &MockCALRecallTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockCALRecallTag_GetAddressString_Call) Run(run func()) *MockCALRecallTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetAddressString_Call) Return(_a0 string) *MockCALRecallTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetAddressString_Call) RunAndReturn(run func() string) *MockCALRecallTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockCALRecallTag) GetArrayInfo() []model.ArrayInfo {
	ret := _m.Called()

	var r0 []model.ArrayInfo
	if rf, ok := ret.Get(0).(func() []model.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ArrayInfo)
		}
	}

	return r0
}

// MockCALRecallTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockCALRecallTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetArrayInfo() *MockCALRecallTag_GetArrayInfo_Call {
	return &MockCALRecallTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockCALRecallTag_GetArrayInfo_Call) Run(run func()) *MockCALRecallTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockCALRecallTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockCALRecallTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetBridgeAddresses provides a mock function with given fields:
func (_m *MockCALRecallTag) GetBridgeAddresses() []readwritemodel.BridgeAddress {
	ret := _m.Called()

	var r0 []readwritemodel.BridgeAddress
	if rf, ok := ret.Get(0).(func() []readwritemodel.BridgeAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]readwritemodel.BridgeAddress)
		}
	}

	return r0
}

// MockCALRecallTag_GetBridgeAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBridgeAddresses'
type MockCALRecallTag_GetBridgeAddresses_Call struct {
	*mock.Call
}

// GetBridgeAddresses is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetBridgeAddresses() *MockCALRecallTag_GetBridgeAddresses_Call {
	return &MockCALRecallTag_GetBridgeAddresses_Call{Call: _e.mock.On("GetBridgeAddresses")}
}

func (_c *MockCALRecallTag_GetBridgeAddresses_Call) Run(run func()) *MockCALRecallTag_GetBridgeAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetBridgeAddresses_Call) Return(_a0 []readwritemodel.BridgeAddress) *MockCALRecallTag_GetBridgeAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetBridgeAddresses_Call) RunAndReturn(run func() []readwritemodel.BridgeAddress) *MockCALRecallTag_GetBridgeAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// GetCount provides a mock function with given fields:
func (_m *MockCALRecallTag) GetCount() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// MockCALRecallTag_GetCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCount'
type MockCALRecallTag_GetCount_Call struct {
	*mock.Call
}

// GetCount is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetCount() *MockCALRecallTag_GetCount_Call {
	return &MockCALRecallTag_GetCount_Call{Call: _e.mock.On("GetCount")}
}

func (_c *MockCALRecallTag_GetCount_Call) Run(run func()) *MockCALRecallTag_GetCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetCount_Call) Return(_a0 uint8) *MockCALRecallTag_GetCount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetCount_Call) RunAndReturn(run func() uint8) *MockCALRecallTag_GetCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetParameter provides a mock function with given fields:
func (_m *MockCALRecallTag) GetParameter() readwritemodel.Parameter {
	ret := _m.Called()

	var r0 readwritemodel.Parameter
	if rf, ok := ret.Get(0).(func() readwritemodel.Parameter); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(readwritemodel.Parameter)
	}

	return r0
}

// MockCALRecallTag_GetParameter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParameter'
type MockCALRecallTag_GetParameter_Call struct {
	*mock.Call
}

// GetParameter is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetParameter() *MockCALRecallTag_GetParameter_Call {
	return &MockCALRecallTag_GetParameter_Call{Call: _e.mock.On("GetParameter")}
}

func (_c *MockCALRecallTag_GetParameter_Call) Run(run func()) *MockCALRecallTag_GetParameter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetParameter_Call) Return(_a0 readwritemodel.Parameter) *MockCALRecallTag_GetParameter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetParameter_Call) RunAndReturn(run func() readwritemodel.Parameter) *MockCALRecallTag_GetParameter_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagType provides a mock function with given fields:
func (_m *MockCALRecallTag) GetTagType() TagType {
	ret := _m.Called()

	var r0 TagType
	if rf, ok := ret.Get(0).(func() TagType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TagType)
	}

	return r0
}

// MockCALRecallTag_GetTagType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagType'
type MockCALRecallTag_GetTagType_Call struct {
	*mock.Call
}

// GetTagType is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetTagType() *MockCALRecallTag_GetTagType_Call {
	return &MockCALRecallTag_GetTagType_Call{Call: _e.mock.On("GetTagType")}
}

func (_c *MockCALRecallTag_GetTagType_Call) Run(run func()) *MockCALRecallTag_GetTagType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetTagType_Call) Return(_a0 TagType) *MockCALRecallTag_GetTagType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetTagType_Call) RunAndReturn(run func() TagType) *MockCALRecallTag_GetTagType_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnitAddress provides a mock function with given fields:
func (_m *MockCALRecallTag) GetUnitAddress() readwritemodel.UnitAddress {
	ret := _m.Called()

	var r0 readwritemodel.UnitAddress
	if rf, ok := ret.Get(0).(func() readwritemodel.UnitAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(readwritemodel.UnitAddress)
		}
	}

	return r0
}

// MockCALRecallTag_GetUnitAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnitAddress'
type MockCALRecallTag_GetUnitAddress_Call struct {
	*mock.Call
}

// GetUnitAddress is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetUnitAddress() *MockCALRecallTag_GetUnitAddress_Call {
	return &MockCALRecallTag_GetUnitAddress_Call{Call: _e.mock.On("GetUnitAddress")}
}

func (_c *MockCALRecallTag_GetUnitAddress_Call) Run(run func()) *MockCALRecallTag_GetUnitAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetUnitAddress_Call) Return(_a0 readwritemodel.UnitAddress) *MockCALRecallTag_GetUnitAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetUnitAddress_Call) RunAndReturn(run func() readwritemodel.UnitAddress) *MockCALRecallTag_GetUnitAddress_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockCALRecallTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockCALRecallTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockCALRecallTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) GetValueType() *MockCALRecallTag_GetValueType_Call {
	return &MockCALRecallTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockCALRecallTag_GetValueType_Call) Run(run func()) *MockCALRecallTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockCALRecallTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockCALRecallTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockCALRecallTag) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCALRecallTag_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockCALRecallTag_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockCALRecallTag_Expecter) String() *MockCALRecallTag_String_Call {
	return &MockCALRecallTag_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockCALRecallTag_String_Call) Run(run func()) *MockCALRecallTag_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCALRecallTag_String_Call) Return(_a0 string) *MockCALRecallTag_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCALRecallTag_String_Call) RunAndReturn(run func() string) *MockCALRecallTag_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCALRecallTag creates a new instance of MockCALRecallTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCALRecallTag(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCALRecallTag {
	mock := &MockCALRecallTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
