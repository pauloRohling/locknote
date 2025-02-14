// Code generated by mockery v2.51.1. DO NOT EDIT.

package mockuser

import (
	context "context"

	email "github.com/pauloRohling/locknote/internal/domain/types/email"
	id "github.com/pauloRohling/locknote/internal/domain/types/id"

	mock "github.com/stretchr/testify/mock"

	user "github.com/pauloRohling/locknote/internal/domain/user"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// DeleteById provides a mock function with given fields: ctx, userId
func (_m *MockRepository) DeleteById(ctx context.Context, userId id.ID) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, id.ID) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockRepository_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - userId id.ID
func (_e *MockRepository_Expecter) DeleteById(ctx interface{}, userId interface{}) *MockRepository_DeleteById_Call {
	return &MockRepository_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, userId)}
}

func (_c *MockRepository_DeleteById_Call) Run(run func(ctx context.Context, userId id.ID)) *MockRepository_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(id.ID))
	})
	return _c
}

func (_c *MockRepository_DeleteById_Call) Return(_a0 error) *MockRepository_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_DeleteById_Call) RunAndReturn(run func(context.Context, id.ID) error) *MockRepository_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: ctx
func (_m *MockRepository) Find(ctx context.Context) (*user.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*user.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *user.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepository_Expecter) Find(ctx interface{}) *MockRepository_Find_Call {
	return &MockRepository_Find_Call{Call: _e.mock.On("Find", ctx)}
}

func (_c *MockRepository_Find_Call) Run(run func(ctx context.Context)) *MockRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepository_Find_Call) Return(_a0 *user.User, _a1 error) *MockRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Find_Call) RunAndReturn(run func(context.Context) (*user.User, error)) *MockRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEmail provides a mock function with given fields: ctx, _a1
func (_m *MockRepository) FindByEmail(ctx context.Context, _a1 email.Email) (*user.User, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, email.Email) (*user.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, email.Email) *user.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, email.Email) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEmail'
type MockRepository_FindByEmail_Call struct {
	*mock.Call
}

// FindByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 email.Email
func (_e *MockRepository_Expecter) FindByEmail(ctx interface{}, _a1 interface{}) *MockRepository_FindByEmail_Call {
	return &MockRepository_FindByEmail_Call{Call: _e.mock.On("FindByEmail", ctx, _a1)}
}

func (_c *MockRepository_FindByEmail_Call) Run(run func(ctx context.Context, _a1 email.Email)) *MockRepository_FindByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(email.Email))
	})
	return _c
}

func (_c *MockRepository_FindByEmail_Call) Return(_a0 *user.User, _a1 error) *MockRepository_FindByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindByEmail_Call) RunAndReturn(run func(context.Context, email.Email) (*user.User, error)) *MockRepository_FindByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *MockRepository) Save(ctx context.Context, _a1 *user.User) (*user.User, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.User) (*user.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.User) *user.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *user.User
func (_e *MockRepository_Expecter) Save(ctx interface{}, _a1 interface{}) *MockRepository_Save_Call {
	return &MockRepository_Save_Call{Call: _e.mock.On("Save", ctx, _a1)}
}

func (_c *MockRepository_Save_Call) Run(run func(ctx context.Context, _a1 *user.User)) *MockRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.User))
	})
	return _c
}

func (_c *MockRepository_Save_Call) Return(_a0 *user.User, _a1 error) *MockRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Save_Call) RunAndReturn(run func(context.Context, *user.User) (*user.User, error)) *MockRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateById provides a mock function with given fields: ctx, _a1
func (_m *MockRepository) UpdateById(ctx context.Context, _a1 *user.User) (*user.User, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateById")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.User) (*user.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.User) *user.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.User) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_UpdateById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateById'
type MockRepository_UpdateById_Call struct {
	*mock.Call
}

// UpdateById is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *user.User
func (_e *MockRepository_Expecter) UpdateById(ctx interface{}, _a1 interface{}) *MockRepository_UpdateById_Call {
	return &MockRepository_UpdateById_Call{Call: _e.mock.On("UpdateById", ctx, _a1)}
}

func (_c *MockRepository_UpdateById_Call) Run(run func(ctx context.Context, _a1 *user.User)) *MockRepository_UpdateById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*user.User))
	})
	return _c
}

func (_c *MockRepository_UpdateById_Call) Return(_a0 *user.User, _a1 error) *MockRepository_UpdateById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_UpdateById_Call) RunAndReturn(run func(context.Context, *user.User) (*user.User, error)) *MockRepository_UpdateById_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
