// Code generated by mockery v2.46.2. DO NOT EDIT.

package logger_mock

import mock "github.com/stretchr/testify/mock"

// MocktestingLogger is an autogenerated mock type for the testingLogger type
type MocktestingLogger struct {
	mock.Mock
}

type MocktestingLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MocktestingLogger) EXPECT() *MocktestingLogger_Expecter {
	return &MocktestingLogger_Expecter{mock: &_m.Mock}
}

// Logf provides a mock function with given fields: format, v
func (_m *MocktestingLogger) Logf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// MocktestingLogger_Logf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logf'
type MocktestingLogger_Logf_Call struct {
	*mock.Call
}

// Logf is a helper method to define mock.On call
//   - format string
//   - v ...interface{}
func (_e *MocktestingLogger_Expecter) Logf(format interface{}, v ...interface{}) *MocktestingLogger_Logf_Call {
	return &MocktestingLogger_Logf_Call{Call: _e.mock.On("Logf",
		append([]interface{}{format}, v...)...)}
}

func (_c *MocktestingLogger_Logf_Call) Run(run func(format string, v ...interface{})) *MocktestingLogger_Logf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MocktestingLogger_Logf_Call) Return() *MocktestingLogger_Logf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MocktestingLogger_Logf_Call) RunAndReturn(run func(string, ...interface{})) *MocktestingLogger_Logf_Call {
	_c.Call.Return(run)
	return _c
}

// NewMocktestingLogger creates a new instance of MocktestingLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMocktestingLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MocktestingLogger {
	mock := &MocktestingLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
