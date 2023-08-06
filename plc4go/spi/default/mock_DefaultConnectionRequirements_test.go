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

package _default

import (
	context "context"

	plc4go "github.com/apache/plc4x/plc4go/pkg/api"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"
)

// MockDefaultConnectionRequirements is an autogenerated mock type for the DefaultConnectionRequirements type
type MockDefaultConnectionRequirements struct {
	mock.Mock
}

type MockDefaultConnectionRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultConnectionRequirements) EXPECT() *MockDefaultConnectionRequirements_Expecter {
	return &MockDefaultConnectionRequirements_Expecter{mock: &_m.Mock}
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockDefaultConnectionRequirements) ConnectWithContext(ctx context.Context) <-chan plc4go.PlcConnectionConnectResult {
	ret := _m.Called(ctx)

	var r0 <-chan plc4go.PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan plc4go.PlcConnectionConnectResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan plc4go.PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockDefaultConnectionRequirements_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockDefaultConnectionRequirements_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDefaultConnectionRequirements_Expecter) ConnectWithContext(ctx interface{}) *MockDefaultConnectionRequirements_ConnectWithContext_Call {
	return &MockDefaultConnectionRequirements_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockDefaultConnectionRequirements_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockDefaultConnectionRequirements_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDefaultConnectionRequirements_ConnectWithContext_Call) Return(_a0 <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnectionRequirements_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnectionRequirements_ConnectWithContext_Call) RunAndReturn(run func(context.Context) <-chan plc4go.PlcConnectionConnectResult) *MockDefaultConnectionRequirements_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields:
func (_m *MockDefaultConnectionRequirements) GetConnection() plc4go.PlcConnection {
	ret := _m.Called()

	var r0 plc4go.PlcConnection
	if rf, ok := ret.Get(0).(func() plc4go.PlcConnection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plc4go.PlcConnection)
		}
	}

	return r0
}

// MockDefaultConnectionRequirements_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockDefaultConnectionRequirements_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
func (_e *MockDefaultConnectionRequirements_Expecter) GetConnection() *MockDefaultConnectionRequirements_GetConnection_Call {
	return &MockDefaultConnectionRequirements_GetConnection_Call{Call: _e.mock.On("GetConnection")}
}

func (_c *MockDefaultConnectionRequirements_GetConnection_Call) Run(run func()) *MockDefaultConnectionRequirements_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnectionRequirements_GetConnection_Call) Return(_a0 plc4go.PlcConnection) *MockDefaultConnectionRequirements_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnectionRequirements_GetConnection_Call) RunAndReturn(run func() plc4go.PlcConnection) *MockDefaultConnectionRequirements_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessageCodec provides a mock function with given fields:
func (_m *MockDefaultConnectionRequirements) GetMessageCodec() spi.MessageCodec {
	ret := _m.Called()

	var r0 spi.MessageCodec
	if rf, ok := ret.Get(0).(func() spi.MessageCodec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.MessageCodec)
		}
	}

	return r0
}

// MockDefaultConnectionRequirements_GetMessageCodec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessageCodec'
type MockDefaultConnectionRequirements_GetMessageCodec_Call struct {
	*mock.Call
}

// GetMessageCodec is a helper method to define mock.On call
func (_e *MockDefaultConnectionRequirements_Expecter) GetMessageCodec() *MockDefaultConnectionRequirements_GetMessageCodec_Call {
	return &MockDefaultConnectionRequirements_GetMessageCodec_Call{Call: _e.mock.On("GetMessageCodec")}
}

func (_c *MockDefaultConnectionRequirements_GetMessageCodec_Call) Run(run func()) *MockDefaultConnectionRequirements_GetMessageCodec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultConnectionRequirements_GetMessageCodec_Call) Return(_a0 spi.MessageCodec) *MockDefaultConnectionRequirements_GetMessageCodec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultConnectionRequirements_GetMessageCodec_Call) RunAndReturn(run func() spi.MessageCodec) *MockDefaultConnectionRequirements_GetMessageCodec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDefaultConnectionRequirements creates a new instance of MockDefaultConnectionRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDefaultConnectionRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDefaultConnectionRequirements {
	mock := &MockDefaultConnectionRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
