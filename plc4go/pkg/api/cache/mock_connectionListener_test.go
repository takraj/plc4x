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

package cache

import mock "github.com/stretchr/testify/mock"

// mockConnectionListener is an autogenerated mock type for the connectionListener type
type mockConnectionListener struct {
	mock.Mock
}

type mockConnectionListener_Expecter struct {
	mock *mock.Mock
}

func (_m *mockConnectionListener) EXPECT() *mockConnectionListener_Expecter {
	return &mockConnectionListener_Expecter{mock: &_m.Mock}
}

// onConnectionEvent provides a mock function with given fields: event
func (_m *mockConnectionListener) onConnectionEvent(event connectionEvent) {
	_m.Called(event)
}

// mockConnectionListener_onConnectionEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'onConnectionEvent'
type mockConnectionListener_onConnectionEvent_Call struct {
	*mock.Call
}

// onConnectionEvent is a helper method to define mock.On call
//   - event connectionEvent
func (_e *mockConnectionListener_Expecter) onConnectionEvent(event interface{}) *mockConnectionListener_onConnectionEvent_Call {
	return &mockConnectionListener_onConnectionEvent_Call{Call: _e.mock.On("onConnectionEvent", event)}
}

func (_c *mockConnectionListener_onConnectionEvent_Call) Run(run func(event connectionEvent)) *mockConnectionListener_onConnectionEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(connectionEvent))
	})
	return _c
}

func (_c *mockConnectionListener_onConnectionEvent_Call) Return() *mockConnectionListener_onConnectionEvent_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockConnectionListener_onConnectionEvent_Call) RunAndReturn(run func(connectionEvent)) *mockConnectionListener_onConnectionEvent_Call {
	_c.Call.Return(run)
	return _c
}

// newMockConnectionListener creates a new instance of mockConnectionListener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockConnectionListener(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockConnectionListener {
	mock := &mockConnectionListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
