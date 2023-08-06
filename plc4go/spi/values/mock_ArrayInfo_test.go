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

package values

import mock "github.com/stretchr/testify/mock"

// MockArrayInfo is an autogenerated mock type for the ArrayInfo type
type MockArrayInfo struct {
	mock.Mock
}

type MockArrayInfo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockArrayInfo) EXPECT() *MockArrayInfo_Expecter {
	return &MockArrayInfo_Expecter{mock: &_m.Mock}
}

// GetLowerBound provides a mock function with given fields:
func (_m *MockArrayInfo) GetLowerBound() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MockArrayInfo_GetLowerBound_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLowerBound'
type MockArrayInfo_GetLowerBound_Call struct {
	*mock.Call
}

// GetLowerBound is a helper method to define mock.On call
func (_e *MockArrayInfo_Expecter) GetLowerBound() *MockArrayInfo_GetLowerBound_Call {
	return &MockArrayInfo_GetLowerBound_Call{Call: _e.mock.On("GetLowerBound")}
}

func (_c *MockArrayInfo_GetLowerBound_Call) Run(run func()) *MockArrayInfo_GetLowerBound_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArrayInfo_GetLowerBound_Call) Return(_a0 uint32) *MockArrayInfo_GetLowerBound_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArrayInfo_GetLowerBound_Call) RunAndReturn(run func() uint32) *MockArrayInfo_GetLowerBound_Call {
	_c.Call.Return(run)
	return _c
}

// GetSize provides a mock function with given fields:
func (_m *MockArrayInfo) GetSize() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MockArrayInfo_GetSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSize'
type MockArrayInfo_GetSize_Call struct {
	*mock.Call
}

// GetSize is a helper method to define mock.On call
func (_e *MockArrayInfo_Expecter) GetSize() *MockArrayInfo_GetSize_Call {
	return &MockArrayInfo_GetSize_Call{Call: _e.mock.On("GetSize")}
}

func (_c *MockArrayInfo_GetSize_Call) Run(run func()) *MockArrayInfo_GetSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArrayInfo_GetSize_Call) Return(_a0 uint32) *MockArrayInfo_GetSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArrayInfo_GetSize_Call) RunAndReturn(run func() uint32) *MockArrayInfo_GetSize_Call {
	_c.Call.Return(run)
	return _c
}

// GetUpperBound provides a mock function with given fields:
func (_m *MockArrayInfo) GetUpperBound() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// MockArrayInfo_GetUpperBound_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUpperBound'
type MockArrayInfo_GetUpperBound_Call struct {
	*mock.Call
}

// GetUpperBound is a helper method to define mock.On call
func (_e *MockArrayInfo_Expecter) GetUpperBound() *MockArrayInfo_GetUpperBound_Call {
	return &MockArrayInfo_GetUpperBound_Call{Call: _e.mock.On("GetUpperBound")}
}

func (_c *MockArrayInfo_GetUpperBound_Call) Run(run func()) *MockArrayInfo_GetUpperBound_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArrayInfo_GetUpperBound_Call) Return(_a0 uint32) *MockArrayInfo_GetUpperBound_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArrayInfo_GetUpperBound_Call) RunAndReturn(run func() uint32) *MockArrayInfo_GetUpperBound_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockArrayInfo) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockArrayInfo_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockArrayInfo_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockArrayInfo_Expecter) String() *MockArrayInfo_String_Call {
	return &MockArrayInfo_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockArrayInfo_String_Call) Run(run func()) *MockArrayInfo_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockArrayInfo_String_Call) Return(_a0 string) *MockArrayInfo_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArrayInfo_String_Call) RunAndReturn(run func() string) *MockArrayInfo_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockArrayInfo creates a new instance of MockArrayInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockArrayInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockArrayInfo {
	mock := &MockArrayInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
