// Code generated by mockery v2.20.2. DO NOT EDIT.

package core

import (
	models "github.com/keldonia/btime.go/models"
	mock "github.com/stretchr/testify/mock"
)

// MockBScheduleUtil is an autogenerated mock type for the BScheduleUtil type
type MockBScheduleUtil struct {
	mock.Mock
}

type MockBScheduleUtil_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBScheduleUtil) EXPECT() *MockBScheduleUtil_Expecter {
	return &MockBScheduleUtil_Expecter{mock: &_m.Mock}
}

// DeleteAppointment provides a mock function with given fields: timeSlotToDelete, scheduleSlot
func (_m *MockBScheduleUtil) DeleteAppointment(timeSlotToDelete *models.Appointment, scheduleSlot string) (*string, error) {
	ret := _m.Called(timeSlotToDelete, scheduleSlot)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Appointment, string) (*string, error)); ok {
		return rf(timeSlotToDelete, scheduleSlot)
	}
	if rf, ok := ret.Get(0).(func(*models.Appointment, string) *string); ok {
		r0 = rf(timeSlotToDelete, scheduleSlot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Appointment, string) error); ok {
		r1 = rf(timeSlotToDelete, scheduleSlot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_DeleteAppointment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAppointment'
type MockBScheduleUtil_DeleteAppointment_Call struct {
	*mock.Call
}

// DeleteAppointment is a helper method to define mock.On call
//   - timeSlotToDelete *models.Appointment
//   - scheduleSlot string
func (_e *MockBScheduleUtil_Expecter) DeleteAppointment(timeSlotToDelete interface{}, scheduleSlot interface{}) *MockBScheduleUtil_DeleteAppointment_Call {
	return &MockBScheduleUtil_DeleteAppointment_Call{Call: _e.mock.On("DeleteAppointment", timeSlotToDelete, scheduleSlot)}
}

func (_c *MockBScheduleUtil_DeleteAppointment_Call) Run(run func(timeSlotToDelete *models.Appointment, scheduleSlot string)) *MockBScheduleUtil_DeleteAppointment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointment_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_DeleteAppointment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointment_Call) RunAndReturn(run func(*models.Appointment, string) (*string, error)) *MockBScheduleUtil_DeleteAppointment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAppointmentBString provides a mock function with given fields: bStringToDelete, scheduleSlot
func (_m *MockBScheduleUtil) DeleteAppointmentBString(bStringToDelete string, scheduleSlot string) (*string, error) {
	ret := _m.Called(bStringToDelete, scheduleSlot)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*string, error)); ok {
		return rf(bStringToDelete, scheduleSlot)
	}
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(bStringToDelete, scheduleSlot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bStringToDelete, scheduleSlot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_DeleteAppointmentBString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAppointmentBString'
type MockBScheduleUtil_DeleteAppointmentBString_Call struct {
	*mock.Call
}

// DeleteAppointmentBString is a helper method to define mock.On call
//   - bStringToDelete string
//   - scheduleSlot string
func (_e *MockBScheduleUtil_Expecter) DeleteAppointmentBString(bStringToDelete interface{}, scheduleSlot interface{}) *MockBScheduleUtil_DeleteAppointmentBString_Call {
	return &MockBScheduleUtil_DeleteAppointmentBString_Call{Call: _e.mock.On("DeleteAppointmentBString", bStringToDelete, scheduleSlot)}
}

func (_c *MockBScheduleUtil_DeleteAppointmentBString_Call) Run(run func(bStringToDelete string, scheduleSlot string)) *MockBScheduleUtil_DeleteAppointmentBString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointmentBString_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_DeleteAppointmentBString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointmentBString_Call) RunAndReturn(run func(string, string) (*string, error)) *MockBScheduleUtil_DeleteAppointmentBString_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAppointmentInterval provides a mock function with given fields: timeSlotBString, scheduleInterval
func (_m *MockBScheduleUtil) DeleteAppointmentInterval(timeSlotBString string, scheduleInterval string) (*string, error) {
	ret := _m.Called(timeSlotBString, scheduleInterval)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*string, error)); ok {
		return rf(timeSlotBString, scheduleInterval)
	}
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(timeSlotBString, scheduleInterval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(timeSlotBString, scheduleInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_DeleteAppointmentInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAppointmentInterval'
type MockBScheduleUtil_DeleteAppointmentInterval_Call struct {
	*mock.Call
}

// DeleteAppointmentInterval is a helper method to define mock.On call
//   - timeSlotBString string
//   - scheduleInterval string
func (_e *MockBScheduleUtil_Expecter) DeleteAppointmentInterval(timeSlotBString interface{}, scheduleInterval interface{}) *MockBScheduleUtil_DeleteAppointmentInterval_Call {
	return &MockBScheduleUtil_DeleteAppointmentInterval_Call{Call: _e.mock.On("DeleteAppointmentInterval", timeSlotBString, scheduleInterval)}
}

func (_c *MockBScheduleUtil_DeleteAppointmentInterval_Call) Run(run func(timeSlotBString string, scheduleInterval string)) *MockBScheduleUtil_DeleteAppointmentInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointmentInterval_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_DeleteAppointmentInterval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_DeleteAppointmentInterval_Call) RunAndReturn(run func(string, string) (*string, error)) *MockBScheduleUtil_DeleteAppointmentInterval_Call {
	_c.Call.Return(run)
	return _c
}

// MergeScheduleBStringWithTest provides a mock function with given fields: timeSlotBString, schedule
func (_m *MockBScheduleUtil) MergeScheduleBStringWithTest(timeSlotBString string, schedule string) (*string, error) {
	ret := _m.Called(timeSlotBString, schedule)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*string, error)); ok {
		return rf(timeSlotBString, schedule)
	}
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(timeSlotBString, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(timeSlotBString, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_MergeScheduleBStringWithTest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeScheduleBStringWithTest'
type MockBScheduleUtil_MergeScheduleBStringWithTest_Call struct {
	*mock.Call
}

// MergeScheduleBStringWithTest is a helper method to define mock.On call
//   - timeSlotBString string
//   - schedule string
func (_e *MockBScheduleUtil_Expecter) MergeScheduleBStringWithTest(timeSlotBString interface{}, schedule interface{}) *MockBScheduleUtil_MergeScheduleBStringWithTest_Call {
	return &MockBScheduleUtil_MergeScheduleBStringWithTest_Call{Call: _e.mock.On("MergeScheduleBStringWithTest", timeSlotBString, schedule)}
}

func (_c *MockBScheduleUtil_MergeScheduleBStringWithTest_Call) Run(run func(timeSlotBString string, schedule string)) *MockBScheduleUtil_MergeScheduleBStringWithTest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringWithTest_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_MergeScheduleBStringWithTest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringWithTest_Call) RunAndReturn(run func(string, string) (*string, error)) *MockBScheduleUtil_MergeScheduleBStringWithTest_Call {
	_c.Call.Return(run)
	return _c
}

// MergeScheduleBStringsWithTest provides a mock function with given fields: timeSlot, schedule
func (_m *MockBScheduleUtil) MergeScheduleBStringsWithTest(timeSlot *models.Appointment, schedule string) (*string, error) {
	ret := _m.Called(timeSlot, schedule)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Appointment, string) (*string, error)); ok {
		return rf(timeSlot, schedule)
	}
	if rf, ok := ret.Get(0).(func(*models.Appointment, string) *string); ok {
		r0 = rf(timeSlot, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Appointment, string) error); ok {
		r1 = rf(timeSlot, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_MergeScheduleBStringsWithTest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeScheduleBStringsWithTest'
type MockBScheduleUtil_MergeScheduleBStringsWithTest_Call struct {
	*mock.Call
}

// MergeScheduleBStringsWithTest is a helper method to define mock.On call
//   - timeSlot *models.Appointment
//   - schedule string
func (_e *MockBScheduleUtil_Expecter) MergeScheduleBStringsWithTest(timeSlot interface{}, schedule interface{}) *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call {
	return &MockBScheduleUtil_MergeScheduleBStringsWithTest_Call{Call: _e.mock.On("MergeScheduleBStringsWithTest", timeSlot, schedule)}
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call) Run(run func(timeSlot *models.Appointment, schedule string)) *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call) RunAndReturn(run func(*models.Appointment, string) (*string, error)) *MockBScheduleUtil_MergeScheduleBStringsWithTest_Call {
	_c.Call.Return(run)
	return _c
}

// MergeScheduleBStringsWithTestBase provides a mock function with given fields: apptBString, schedule
func (_m *MockBScheduleUtil) MergeScheduleBStringsWithTestBase(apptBString string, schedule string) (*string, error) {
	ret := _m.Called(apptBString, schedule)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*string, error)); ok {
		return rf(apptBString, schedule)
	}
	if rf, ok := ret.Get(0).(func(string, string) *string); ok {
		r0 = rf(apptBString, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(apptBString, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeScheduleBStringsWithTestBase'
type MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call struct {
	*mock.Call
}

// MergeScheduleBStringsWithTestBase is a helper method to define mock.On call
//   - apptBString string
//   - schedule string
func (_e *MockBScheduleUtil_Expecter) MergeScheduleBStringsWithTestBase(apptBString interface{}, schedule interface{}) *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call {
	return &MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call{Call: _e.mock.On("MergeScheduleBStringsWithTestBase", apptBString, schedule)}
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call) Run(run func(apptBString string, schedule string)) *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call) RunAndReturn(run func(string, string) (*string, error)) *MockBScheduleUtil_MergeScheduleBStringsWithTestBase_Call {
	_c.Call.Return(run)
	return _c
}

// ModifyScheduleAndBooking provides a mock function with given fields: scheduleBStringToModify, scheduleBStringToTest, appt
func (_m *MockBScheduleUtil) ModifyScheduleAndBooking(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	ret := _m.Called(scheduleBStringToModify, scheduleBStringToTest, appt)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*string, error)); ok {
		return rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *string); ok {
		r0 = rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_ModifyScheduleAndBooking_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ModifyScheduleAndBooking'
type MockBScheduleUtil_ModifyScheduleAndBooking_Call struct {
	*mock.Call
}

// ModifyScheduleAndBooking is a helper method to define mock.On call
//   - scheduleBStringToModify string
//   - scheduleBStringToTest string
//   - appt string
func (_e *MockBScheduleUtil_Expecter) ModifyScheduleAndBooking(scheduleBStringToModify interface{}, scheduleBStringToTest interface{}, appt interface{}) *MockBScheduleUtil_ModifyScheduleAndBooking_Call {
	return &MockBScheduleUtil_ModifyScheduleAndBooking_Call{Call: _e.mock.On("ModifyScheduleAndBooking", scheduleBStringToModify, scheduleBStringToTest, appt)}
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBooking_Call) Run(run func(scheduleBStringToModify string, scheduleBStringToTest string, appt string)) *MockBScheduleUtil_ModifyScheduleAndBooking_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBooking_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_ModifyScheduleAndBooking_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBooking_Call) RunAndReturn(run func(string, string, string) (*string, error)) *MockBScheduleUtil_ModifyScheduleAndBooking_Call {
	_c.Call.Return(run)
	return _c
}

// ModifyScheduleAndBookingInterval provides a mock function with given fields: scheduleBStringToModify, scheduleBStringToTest, appt
func (_m *MockBScheduleUtil) ModifyScheduleAndBookingInterval(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	ret := _m.Called(scheduleBStringToModify, scheduleBStringToTest, appt)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*string, error)); ok {
		return rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *string); ok {
		r0 = rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(scheduleBStringToModify, scheduleBStringToTest, appt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ModifyScheduleAndBookingInterval'
type MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call struct {
	*mock.Call
}

// ModifyScheduleAndBookingInterval is a helper method to define mock.On call
//   - scheduleBStringToModify string
//   - scheduleBStringToTest string
//   - appt string
func (_e *MockBScheduleUtil_Expecter) ModifyScheduleAndBookingInterval(scheduleBStringToModify interface{}, scheduleBStringToTest interface{}, appt interface{}) *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call {
	return &MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call{Call: _e.mock.On("ModifyScheduleAndBookingInterval", scheduleBStringToModify, scheduleBStringToTest, appt)}
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call) Run(run func(scheduleBStringToModify string, scheduleBStringToTest string, appt string)) *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call) Return(_a0 *string, _a1 error) *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call) RunAndReturn(run func(string, string, string) (*string, error)) *MockBScheduleUtil_ModifyScheduleAndBookingInterval_Call {
	_c.Call.Return(run)
	return _c
}

// TestViabilityAndCompute provides a mock function with given fields: binary1, binary2
func (_m *MockBScheduleUtil) TestViabilityAndCompute(binary1 int64, binary2 int64) (*int64, error) {
	ret := _m.Called(binary1, binary2)

	var r0 *int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (*int64, error)); ok {
		return rf(binary1, binary2)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) *int64); ok {
		r0 = rf(binary1, binary2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(binary1, binary2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduleUtil_TestViabilityAndCompute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestViabilityAndCompute'
type MockBScheduleUtil_TestViabilityAndCompute_Call struct {
	*mock.Call
}

// TestViabilityAndCompute is a helper method to define mock.On call
//   - binary1 int64
//   - binary2 int64
func (_e *MockBScheduleUtil_Expecter) TestViabilityAndCompute(binary1 interface{}, binary2 interface{}) *MockBScheduleUtil_TestViabilityAndCompute_Call {
	return &MockBScheduleUtil_TestViabilityAndCompute_Call{Call: _e.mock.On("TestViabilityAndCompute", binary1, binary2)}
}

func (_c *MockBScheduleUtil_TestViabilityAndCompute_Call) Run(run func(binary1 int64, binary2 int64)) *MockBScheduleUtil_TestViabilityAndCompute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *MockBScheduleUtil_TestViabilityAndCompute_Call) Return(_a0 *int64, _a1 error) *MockBScheduleUtil_TestViabilityAndCompute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduleUtil_TestViabilityAndCompute_Call) RunAndReturn(run func(int64, int64) (*int64, error)) *MockBScheduleUtil_TestViabilityAndCompute_Call {
	_c.Call.Return(run)
	return _c
}

// ValidDeletion provides a mock function with given fields: baseNumber, toDeleteNumber
func (_m *MockBScheduleUtil) ValidDeletion(baseNumber int64, toDeleteNumber int64) bool {
	ret := _m.Called(baseNumber, toDeleteNumber)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, int64) bool); ok {
		r0 = rf(baseNumber, toDeleteNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockBScheduleUtil_ValidDeletion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidDeletion'
type MockBScheduleUtil_ValidDeletion_Call struct {
	*mock.Call
}

// ValidDeletion is a helper method to define mock.On call
//   - baseNumber int64
//   - toDeleteNumber int64
func (_e *MockBScheduleUtil_Expecter) ValidDeletion(baseNumber interface{}, toDeleteNumber interface{}) *MockBScheduleUtil_ValidDeletion_Call {
	return &MockBScheduleUtil_ValidDeletion_Call{Call: _e.mock.On("ValidDeletion", baseNumber, toDeleteNumber)}
}

func (_c *MockBScheduleUtil_ValidDeletion_Call) Run(run func(baseNumber int64, toDeleteNumber int64)) *MockBScheduleUtil_ValidDeletion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *MockBScheduleUtil_ValidDeletion_Call) Return(_a0 bool) *MockBScheduleUtil_ValidDeletion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBScheduleUtil_ValidDeletion_Call) RunAndReturn(run func(int64, int64) bool) *MockBScheduleUtil_ValidDeletion_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBScheduleUtil interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBScheduleUtil creates a new instance of MockBScheduleUtil. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBScheduleUtil(t mockConstructorTestingTNewMockBScheduleUtil) *MockBScheduleUtil {
	mock := &MockBScheduleUtil{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}