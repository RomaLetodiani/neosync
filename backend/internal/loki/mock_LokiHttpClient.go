// Code generated by mockery. DO NOT EDIT.

package loki

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockLokiHttpClient is an autogenerated mock type for the LokiHttpClient type
type MockLokiHttpClient struct {
	mock.Mock
}

type MockLokiHttpClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLokiHttpClient) EXPECT() *MockLokiHttpClient_Expecter {
	return &MockLokiHttpClient_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: req
func (_m *MockLokiHttpClient) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*http.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLokiHttpClient_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type MockLokiHttpClient_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - req *http.Request
func (_e *MockLokiHttpClient_Expecter) Do(req interface{}) *MockLokiHttpClient_Do_Call {
	return &MockLokiHttpClient_Do_Call{Call: _e.mock.On("Do", req)}
}

func (_c *MockLokiHttpClient_Do_Call) Run(run func(req *http.Request)) *MockLokiHttpClient_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request))
	})
	return _c
}

func (_c *MockLokiHttpClient_Do_Call) Return(_a0 *http.Response, _a1 error) *MockLokiHttpClient_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLokiHttpClient_Do_Call) RunAndReturn(run func(*http.Request) (*http.Response, error)) *MockLokiHttpClient_Do_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLokiHttpClient creates a new instance of MockLokiHttpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLokiHttpClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLokiHttpClient {
	mock := &MockLokiHttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
