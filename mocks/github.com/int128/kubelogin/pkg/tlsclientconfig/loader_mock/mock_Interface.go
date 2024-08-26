// Code generated by mockery v2.44.2. DO NOT EDIT.

package loader_mock

import (
	tls "crypto/tls"

	mock "github.com/stretchr/testify/mock"

	tlsclientconfig "github.com/int128/kubelogin/pkg/tlsclientconfig"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Load provides a mock function with given fields: config
func (_m *MockInterface) Load(config tlsclientconfig.Config) (*tls.Config, error) {
	ret := _m.Called(config)

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 *tls.Config
	var r1 error
	if rf, ok := ret.Get(0).(func(tlsclientconfig.Config) (*tls.Config, error)); ok {
		return rf(config)
	}
	if rf, ok := ret.Get(0).(func(tlsclientconfig.Config) *tls.Config); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tls.Config)
		}
	}

	if rf, ok := ret.Get(1).(func(tlsclientconfig.Config) error); ok {
		r1 = rf(config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type MockInterface_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - config tlsclientconfig.Config
func (_e *MockInterface_Expecter) Load(config interface{}) *MockInterface_Load_Call {
	return &MockInterface_Load_Call{Call: _e.mock.On("Load", config)}
}

func (_c *MockInterface_Load_Call) Run(run func(config tlsclientconfig.Config)) *MockInterface_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(tlsclientconfig.Config))
	})
	return _c
}

func (_c *MockInterface_Load_Call) Return(_a0 *tls.Config, _a1 error) *MockInterface_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Load_Call) RunAndReturn(run func(tlsclientconfig.Config) (*tls.Config, error)) *MockInterface_Load_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
