// Code generated by mockery v2.51.1. DO NOT EDIT.

package mockuser

import (
	context "context"

	user "github.com/pauloRohling/locknote/internal/domain/user"
	mock "github.com/stretchr/testify/mock"
)

// MockFactory is an autogenerated mock type for the Factory type
type MockFactory struct {
	mock.Mock
}

type MockFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFactory) EXPECT() *MockFactory_Expecter {
	return &MockFactory_Expecter{mock: &_m.Mock}
}

// New provides a mock function with given fields: ctx, params
func (_m *MockFactory) New(ctx context.Context, params user.NewParams) (*user.User, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, user.NewParams) (*user.User, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, user.NewParams) *user.User); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, user.NewParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFactory_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type MockFactory_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - ctx context.Context
//   - params user.NewParams
func (_e *MockFactory_Expecter) New(ctx interface{}, params interface{}) *MockFactory_New_Call {
	return &MockFactory_New_Call{Call: _e.mock.On("New", ctx, params)}
}

func (_c *MockFactory_New_Call) Run(run func(ctx context.Context, params user.NewParams)) *MockFactory_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(user.NewParams))
	})
	return _c
}

func (_c *MockFactory_New_Call) Return(_a0 *user.User, _a1 error) *MockFactory_New_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFactory_New_Call) RunAndReturn(run func(context.Context, user.NewParams) (*user.User, error)) *MockFactory_New_Call {
	_c.Call.Return(run)
	return _c
}

// Parse provides a mock function with given fields: params
func (_m *MockFactory) Parse(params user.ParseParams) (*user.User, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for Parse")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(user.ParseParams) (*user.User, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(user.ParseParams) *user.User); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(user.ParseParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFactory_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type MockFactory_Parse_Call struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - params user.ParseParams
func (_e *MockFactory_Expecter) Parse(params interface{}) *MockFactory_Parse_Call {
	return &MockFactory_Parse_Call{Call: _e.mock.On("Parse", params)}
}

func (_c *MockFactory_Parse_Call) Run(run func(params user.ParseParams)) *MockFactory_Parse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.ParseParams))
	})
	return _c
}

func (_c *MockFactory_Parse_Call) Return(_a0 *user.User, _a1 error) *MockFactory_Parse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFactory_Parse_Call) RunAndReturn(run func(user.ParseParams) (*user.User, error)) *MockFactory_Parse_Call {
	_c.Call.Return(run)
	return _c
}

// ParseWithEncryptedPassword provides a mock function with given fields: params
func (_m *MockFactory) ParseWithEncryptedPassword(params user.ParseParams) (*user.User, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for ParseWithEncryptedPassword")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(user.ParseParams) (*user.User, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(user.ParseParams) *user.User); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(user.ParseParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFactory_ParseWithEncryptedPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParseWithEncryptedPassword'
type MockFactory_ParseWithEncryptedPassword_Call struct {
	*mock.Call
}

// ParseWithEncryptedPassword is a helper method to define mock.On call
//   - params user.ParseParams
func (_e *MockFactory_Expecter) ParseWithEncryptedPassword(params interface{}) *MockFactory_ParseWithEncryptedPassword_Call {
	return &MockFactory_ParseWithEncryptedPassword_Call{Call: _e.mock.On("ParseWithEncryptedPassword", params)}
}

func (_c *MockFactory_ParseWithEncryptedPassword_Call) Run(run func(params user.ParseParams)) *MockFactory_ParseWithEncryptedPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.ParseParams))
	})
	return _c
}

func (_c *MockFactory_ParseWithEncryptedPassword_Call) Return(_a0 *user.User, _a1 error) *MockFactory_ParseWithEncryptedPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFactory_ParseWithEncryptedPassword_Call) RunAndReturn(run func(user.ParseParams) (*user.User, error)) *MockFactory_ParseWithEncryptedPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFactory creates a new instance of MockFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFactory {
	mock := &MockFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
