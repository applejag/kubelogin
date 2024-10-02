// Code generated by mockery v2.46.2. DO NOT EDIT.

package reader_mock

import mock "github.com/stretchr/testify/mock"

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

// ReadPassword provides a mock function with given fields: prompt
func (_m *MockInterface) ReadPassword(prompt string) (string, error) {
	ret := _m.Called(prompt)

	if len(ret) == 0 {
		panic("no return value specified for ReadPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(prompt)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ReadPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadPassword'
type MockInterface_ReadPassword_Call struct {
	*mock.Call
}

// ReadPassword is a helper method to define mock.On call
//   - prompt string
func (_e *MockInterface_Expecter) ReadPassword(prompt interface{}) *MockInterface_ReadPassword_Call {
	return &MockInterface_ReadPassword_Call{Call: _e.mock.On("ReadPassword", prompt)}
}

func (_c *MockInterface_ReadPassword_Call) Run(run func(prompt string)) *MockInterface_ReadPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_ReadPassword_Call) Return(_a0 string, _a1 error) *MockInterface_ReadPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ReadPassword_Call) RunAndReturn(run func(string) (string, error)) *MockInterface_ReadPassword_Call {
	_c.Call.Return(run)
	return _c
}

// ReadString provides a mock function with given fields: prompt
func (_m *MockInterface) ReadString(prompt string) (string, error) {
	ret := _m.Called(prompt)

	if len(ret) == 0 {
		panic("no return value specified for ReadString")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(prompt)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prompt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prompt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_ReadString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadString'
type MockInterface_ReadString_Call struct {
	*mock.Call
}

// ReadString is a helper method to define mock.On call
//   - prompt string
func (_e *MockInterface_Expecter) ReadString(prompt interface{}) *MockInterface_ReadString_Call {
	return &MockInterface_ReadString_Call{Call: _e.mock.On("ReadString", prompt)}
}

func (_c *MockInterface_ReadString_Call) Run(run func(prompt string)) *MockInterface_ReadString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_ReadString_Call) Return(_a0 string, _a1 error) *MockInterface_ReadString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_ReadString_Call) RunAndReturn(run func(string) (string, error)) *MockInterface_ReadString_Call {
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
