// Code generated by mockery. DO NOT EDIT.

package migration

import (
	migration "github.com/goravel/framework/contracts/database/migration"
	mock "github.com/stretchr/testify/mock"
)

// ColumnDefinition is an autogenerated mock type for the ColumnDefinition type
type ColumnDefinition struct {
	mock.Mock
}

type ColumnDefinition_Expecter struct {
	mock *mock.Mock
}

func (_m *ColumnDefinition) EXPECT() *ColumnDefinition_Expecter {
	return &ColumnDefinition_Expecter{mock: &_m.Mock}
}

// AutoIncrement provides a mock function with given fields:
func (_m *ColumnDefinition) AutoIncrement() migration.ColumnDefinition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AutoIncrement")
	}

	var r0 migration.ColumnDefinition
	if rf, ok := ret.Get(0).(func() migration.ColumnDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(migration.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_AutoIncrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutoIncrement'
type ColumnDefinition_AutoIncrement_Call struct {
	*mock.Call
}

// AutoIncrement is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) AutoIncrement() *ColumnDefinition_AutoIncrement_Call {
	return &ColumnDefinition_AutoIncrement_Call{Call: _e.mock.On("AutoIncrement")}
}

func (_c *ColumnDefinition_AutoIncrement_Call) Run(run func()) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_AutoIncrement_Call) Return(_a0 migration.ColumnDefinition) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_AutoIncrement_Call) RunAndReturn(run func() migration.ColumnDefinition) *ColumnDefinition_AutoIncrement_Call {
	_c.Call.Return(run)
	return _c
}

// GetAutoIncrement provides a mock function with given fields:
func (_m *ColumnDefinition) GetAutoIncrement() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAutoIncrement")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_GetAutoIncrement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAutoIncrement'
type ColumnDefinition_GetAutoIncrement_Call struct {
	*mock.Call
}

// GetAutoIncrement is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetAutoIncrement() *ColumnDefinition_GetAutoIncrement_Call {
	return &ColumnDefinition_GetAutoIncrement_Call{Call: _e.mock.On("GetAutoIncrement")}
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) Run(run func()) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) Return(_a0 bool) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetAutoIncrement_Call) RunAndReturn(run func() bool) *ColumnDefinition_GetAutoIncrement_Call {
	_c.Call.Return(run)
	return _c
}

// GetChange provides a mock function with given fields:
func (_m *ColumnDefinition) GetChange() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetChange")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_GetChange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChange'
type ColumnDefinition_GetChange_Call struct {
	*mock.Call
}

// GetChange is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetChange() *ColumnDefinition_GetChange_Call {
	return &ColumnDefinition_GetChange_Call{Call: _e.mock.On("GetChange")}
}

func (_c *ColumnDefinition_GetChange_Call) Run(run func()) *ColumnDefinition_GetChange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetChange_Call) Return(_a0 bool) *ColumnDefinition_GetChange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetChange_Call) RunAndReturn(run func() bool) *ColumnDefinition_GetChange_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefault provides a mock function with given fields:
func (_m *ColumnDefinition) GetDefault() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDefault")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ColumnDefinition_GetDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefault'
type ColumnDefinition_GetDefault_Call struct {
	*mock.Call
}

// GetDefault is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetDefault() *ColumnDefinition_GetDefault_Call {
	return &ColumnDefinition_GetDefault_Call{Call: _e.mock.On("GetDefault")}
}

func (_c *ColumnDefinition_GetDefault_Call) Run(run func()) *ColumnDefinition_GetDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetDefault_Call) Return(_a0 interface{}) *ColumnDefinition_GetDefault_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetDefault_Call) RunAndReturn(run func() interface{}) *ColumnDefinition_GetDefault_Call {
	_c.Call.Return(run)
	return _c
}

// GetLength provides a mock function with given fields:
func (_m *ColumnDefinition) GetLength() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLength")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ColumnDefinition_GetLength_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLength'
type ColumnDefinition_GetLength_Call struct {
	*mock.Call
}

// GetLength is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetLength() *ColumnDefinition_GetLength_Call {
	return &ColumnDefinition_GetLength_Call{Call: _e.mock.On("GetLength")}
}

func (_c *ColumnDefinition_GetLength_Call) Run(run func()) *ColumnDefinition_GetLength_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetLength_Call) Return(_a0 int) *ColumnDefinition_GetLength_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetLength_Call) RunAndReturn(run func() int) *ColumnDefinition_GetLength_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *ColumnDefinition) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnDefinition_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type ColumnDefinition_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetName() *ColumnDefinition_GetName_Call {
	return &ColumnDefinition_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *ColumnDefinition_GetName_Call) Run(run func()) *ColumnDefinition_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetName_Call) Return(_a0 string) *ColumnDefinition_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetName_Call) RunAndReturn(run func() string) *ColumnDefinition_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetNullable provides a mock function with given fields:
func (_m *ColumnDefinition) GetNullable() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNullable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ColumnDefinition_GetNullable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNullable'
type ColumnDefinition_GetNullable_Call struct {
	*mock.Call
}

// GetNullable is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetNullable() *ColumnDefinition_GetNullable_Call {
	return &ColumnDefinition_GetNullable_Call{Call: _e.mock.On("GetNullable")}
}

func (_c *ColumnDefinition_GetNullable_Call) Run(run func()) *ColumnDefinition_GetNullable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetNullable_Call) Return(_a0 bool) *ColumnDefinition_GetNullable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetNullable_Call) RunAndReturn(run func() bool) *ColumnDefinition_GetNullable_Call {
	_c.Call.Return(run)
	return _c
}

// GetType provides a mock function with given fields:
func (_m *ColumnDefinition) GetType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ColumnDefinition_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type ColumnDefinition_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) GetType() *ColumnDefinition_GetType_Call {
	return &ColumnDefinition_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *ColumnDefinition_GetType_Call) Run(run func()) *ColumnDefinition_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_GetType_Call) Return(_a0 string) *ColumnDefinition_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_GetType_Call) RunAndReturn(run func() string) *ColumnDefinition_GetType_Call {
	_c.Call.Return(run)
	return _c
}

// Unsigned provides a mock function with given fields:
func (_m *ColumnDefinition) Unsigned() migration.ColumnDefinition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Unsigned")
	}

	var r0 migration.ColumnDefinition
	if rf, ok := ret.Get(0).(func() migration.ColumnDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(migration.ColumnDefinition)
		}
	}

	return r0
}

// ColumnDefinition_Unsigned_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unsigned'
type ColumnDefinition_Unsigned_Call struct {
	*mock.Call
}

// Unsigned is a helper method to define mock.On call
func (_e *ColumnDefinition_Expecter) Unsigned() *ColumnDefinition_Unsigned_Call {
	return &ColumnDefinition_Unsigned_Call{Call: _e.mock.On("Unsigned")}
}

func (_c *ColumnDefinition_Unsigned_Call) Run(run func()) *ColumnDefinition_Unsigned_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ColumnDefinition_Unsigned_Call) Return(_a0 migration.ColumnDefinition) *ColumnDefinition_Unsigned_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ColumnDefinition_Unsigned_Call) RunAndReturn(run func() migration.ColumnDefinition) *ColumnDefinition_Unsigned_Call {
	_c.Call.Return(run)
	return _c
}

// NewColumnDefinition creates a new instance of ColumnDefinition. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewColumnDefinition(t interface {
	mock.TestingT
	Cleanup(func())
}) *ColumnDefinition {
	mock := &ColumnDefinition{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}