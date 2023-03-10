// Code generated by mockery v2.20.2. DO NOT EDIT.

package core

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockBPointerCalculator is an autogenerated mock type for the BPointerCalculator type
type MockBPointerCalculator struct {
	mock.Mock
}

type MockBPointerCalculator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBPointerCalculator) EXPECT() *MockBPointerCalculator_Expecter {
	return &MockBPointerCalculator_Expecter{mock: &_m.Mock}
}

// FindBPointer provides a mock function with given fields: date
func (_m *MockBPointerCalculator) FindBPointer(date *time.Time) int {
	ret := _m.Called(date)

	var r0 int
	if rf, ok := ret.Get(0).(func(*time.Time) int); ok {
		r0 = rf(date)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockBPointerCalculator_FindBPointer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBPointer'
type MockBPointerCalculator_FindBPointer_Call struct {
	*mock.Call
}

// FindBPointer is a helper method to define mock.On call
//   - date *time.Time
func (_e *MockBPointerCalculator_Expecter) FindBPointer(date interface{}) *MockBPointerCalculator_FindBPointer_Call {
	return &MockBPointerCalculator_FindBPointer_Call{Call: _e.mock.On("FindBPointer", date)}
}

func (_c *MockBPointerCalculator_FindBPointer_Call) Run(run func(date *time.Time)) *MockBPointerCalculator_FindBPointer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*time.Time))
	})
	return _c
}

func (_c *MockBPointerCalculator_FindBPointer_Call) Return(_a0 int) *MockBPointerCalculator_FindBPointer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBPointerCalculator_FindBPointer_Call) RunAndReturn(run func(*time.Time) int) *MockBPointerCalculator_FindBPointer_Call {
	_c.Call.Return(run)
	return _c
}

// FindBPointerIncludingDay provides a mock function with given fields: date
func (_m *MockBPointerCalculator) FindBPointerIncludingDay(date *time.Time) int {
	ret := _m.Called(date)

	var r0 int
	if rf, ok := ret.Get(0).(func(*time.Time) int); ok {
		r0 = rf(date)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockBPointerCalculator_FindBPointerIncludingDay_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBPointerIncludingDay'
type MockBPointerCalculator_FindBPointerIncludingDay_Call struct {
	*mock.Call
}

// FindBPointerIncludingDay is a helper method to define mock.On call
//   - date *time.Time
func (_e *MockBPointerCalculator_Expecter) FindBPointerIncludingDay(date interface{}) *MockBPointerCalculator_FindBPointerIncludingDay_Call {
	return &MockBPointerCalculator_FindBPointerIncludingDay_Call{Call: _e.mock.On("FindBPointerIncludingDay", date)}
}

func (_c *MockBPointerCalculator_FindBPointerIncludingDay_Call) Run(run func(date *time.Time)) *MockBPointerCalculator_FindBPointerIncludingDay_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*time.Time))
	})
	return _c
}

func (_c *MockBPointerCalculator_FindBPointerIncludingDay_Call) Return(_a0 int) *MockBPointerCalculator_FindBPointerIncludingDay_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBPointerCalculator_FindBPointerIncludingDay_Call) RunAndReturn(run func(*time.Time) int) *MockBPointerCalculator_FindBPointerIncludingDay_Call {
	_c.Call.Return(run)
	return _c
}

// FindBPointerModiferForDayOfWeek provides a mock function with given fields: date
func (_m *MockBPointerCalculator) FindBPointerModiferForDayOfWeek(date *time.Time) int {
	ret := _m.Called(date)

	var r0 int
	if rf, ok := ret.Get(0).(func(*time.Time) int); ok {
		r0 = rf(date)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBPointerModiferForDayOfWeek'
type MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call struct {
	*mock.Call
}

// FindBPointerModiferForDayOfWeek is a helper method to define mock.On call
//   - date *time.Time
func (_e *MockBPointerCalculator_Expecter) FindBPointerModiferForDayOfWeek(date interface{}) *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call {
	return &MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call{Call: _e.mock.On("FindBPointerModiferForDayOfWeek", date)}
}

func (_c *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call) Run(run func(date *time.Time)) *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*time.Time))
	})
	return _c
}

func (_c *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call) Return(_a0 int) *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call) RunAndReturn(run func(*time.Time) int) *MockBPointerCalculator_FindBPointerModiferForDayOfWeek_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBPointerCalculator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBPointerCalculator creates a new instance of MockBPointerCalculator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBPointerCalculator(t mockConstructorTestingTNewMockBPointerCalculator) *MockBPointerCalculator {
	mock := &MockBPointerCalculator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}