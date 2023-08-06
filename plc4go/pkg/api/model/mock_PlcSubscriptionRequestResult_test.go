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

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcSubscriptionRequestResult is an autogenerated mock type for the PlcSubscriptionRequestResult type
type MockPlcSubscriptionRequestResult struct {
	mock.Mock
}

type MockPlcSubscriptionRequestResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriptionRequestResult) EXPECT() *MockPlcSubscriptionRequestResult_Expecter {
	return &MockPlcSubscriptionRequestResult_Expecter{mock: &_m.Mock}
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequestResult) GetErr() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcSubscriptionRequestResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcSubscriptionRequestResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequestResult_Expecter) GetErr() *MockPlcSubscriptionRequestResult_GetErr_Call {
	return &MockPlcSubscriptionRequestResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcSubscriptionRequestResult_GetErr_Call) Run(run func()) *MockPlcSubscriptionRequestResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetErr_Call) Return(_a0 error) *MockPlcSubscriptionRequestResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcSubscriptionRequestResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// GetRequest provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequestResult) GetRequest() PlcSubscriptionRequest {
	ret := _m.Called()

	var r0 PlcSubscriptionRequest
	if rf, ok := ret.Get(0).(func() PlcSubscriptionRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcSubscriptionRequest)
		}
	}

	return r0
}

// MockPlcSubscriptionRequestResult_GetRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRequest'
type MockPlcSubscriptionRequestResult_GetRequest_Call struct {
	*mock.Call
}

// GetRequest is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequestResult_Expecter) GetRequest() *MockPlcSubscriptionRequestResult_GetRequest_Call {
	return &MockPlcSubscriptionRequestResult_GetRequest_Call{Call: _e.mock.On("GetRequest")}
}

func (_c *MockPlcSubscriptionRequestResult_GetRequest_Call) Run(run func()) *MockPlcSubscriptionRequestResult_GetRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetRequest_Call) Return(_a0 PlcSubscriptionRequest) *MockPlcSubscriptionRequestResult_GetRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetRequest_Call) RunAndReturn(run func() PlcSubscriptionRequest) *MockPlcSubscriptionRequestResult_GetRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetResponse provides a mock function with given fields:
func (_m *MockPlcSubscriptionRequestResult) GetResponse() PlcSubscriptionResponse {
	ret := _m.Called()

	var r0 PlcSubscriptionResponse
	if rf, ok := ret.Get(0).(func() PlcSubscriptionResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcSubscriptionResponse)
		}
	}

	return r0
}

// MockPlcSubscriptionRequestResult_GetResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResponse'
type MockPlcSubscriptionRequestResult_GetResponse_Call struct {
	*mock.Call
}

// GetResponse is a helper method to define mock.On call
func (_e *MockPlcSubscriptionRequestResult_Expecter) GetResponse() *MockPlcSubscriptionRequestResult_GetResponse_Call {
	return &MockPlcSubscriptionRequestResult_GetResponse_Call{Call: _e.mock.On("GetResponse")}
}

func (_c *MockPlcSubscriptionRequestResult_GetResponse_Call) Run(run func()) *MockPlcSubscriptionRequestResult_GetResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetResponse_Call) Return(_a0 PlcSubscriptionResponse) *MockPlcSubscriptionRequestResult_GetResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriptionRequestResult_GetResponse_Call) RunAndReturn(run func() PlcSubscriptionResponse) *MockPlcSubscriptionRequestResult_GetResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcSubscriptionRequestResult creates a new instance of MockPlcSubscriptionRequestResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcSubscriptionRequestResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcSubscriptionRequestResult {
	mock := &MockPlcSubscriptionRequestResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
