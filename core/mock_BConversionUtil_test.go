// Code generated by mockery v2.20.2. DO NOT EDIT.

package core

import (
	models "github.com/keldonia/btime.go/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockBConversionUtil is an autogenerated mock type for the BConversionUtil type
type MockBConversionUtil struct {
	mock.Mock
}

type MockBConversionUtil_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBConversionUtil) EXPECT() *MockBConversionUtil_Expecter {
	return &MockBConversionUtil_Expecter{mock: &_m.Mock}
}

// CalculateDate provides a mock function with given fields: timePointerIndex, baseDate, end
func (_m *MockBConversionUtil) CalculateDate(timePointerIndex int, baseDate *time.Time, end bool) *time.Time {
	ret := _m.Called(timePointerIndex, baseDate, end)

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func(int, *time.Time, bool) *time.Time); ok {
		r0 = rf(timePointerIndex, baseDate, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	return r0
}

// MockBConversionUtil_CalculateDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateDate'
type MockBConversionUtil_CalculateDate_Call struct {
	*mock.Call
}

// CalculateDate is a helper method to define mock.On call
//   - timePointerIndex int
//   - baseDate *time.Time
//   - end bool
func (_e *MockBConversionUtil_Expecter) CalculateDate(timePointerIndex interface{}, baseDate interface{}, end interface{}) *MockBConversionUtil_CalculateDate_Call {
	return &MockBConversionUtil_CalculateDate_Call{Call: _e.mock.On("CalculateDate", timePointerIndex, baseDate, end)}
}

func (_c *MockBConversionUtil_CalculateDate_Call) Run(run func(timePointerIndex int, baseDate *time.Time, end bool)) *MockBConversionUtil_CalculateDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(*time.Time), args[2].(bool))
	})
	return _c
}

func (_c *MockBConversionUtil_CalculateDate_Call) Return(_a0 *time.Time) *MockBConversionUtil_CalculateDate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBConversionUtil_CalculateDate_Call) RunAndReturn(run func(int, *time.Time, bool) *time.Time) *MockBConversionUtil_CalculateDate_Call {
	_c.Call.Return(run)
	return _c
}

// ConvertScheduleToAppointmentSchedule provides a mock function with given fields: schedule, availability
func (_m *MockBConversionUtil) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule {
	ret := _m.Called(schedule, availability)

	var r0 *models.AppointmentSchedule
	if rf, ok := ret.Get(0).(func(*models.Schedule, []string) *models.AppointmentSchedule); ok {
		r0 = rf(schedule, availability)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AppointmentSchedule)
		}
	}

	return r0
}

// MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertScheduleToAppointmentSchedule'
type MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call struct {
	*mock.Call
}

// ConvertScheduleToAppointmentSchedule is a helper method to define mock.On call
//   - schedule *models.Schedule
//   - availability []string
func (_e *MockBConversionUtil_Expecter) ConvertScheduleToAppointmentSchedule(schedule interface{}, availability interface{}) *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call {
	return &MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call{Call: _e.mock.On("ConvertScheduleToAppointmentSchedule", schedule, availability)}
}

func (_c *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call) Run(run func(schedule *models.Schedule, availability []string)) *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Schedule), args[1].([]string))
	})
	return _c
}

func (_c *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call) Return(_a0 *models.AppointmentSchedule) *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call) RunAndReturn(run func(*models.Schedule, []string) *models.AppointmentSchedule) *MockBConversionUtil_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// ConvertTimeSlotsStringToAppointments provides a mock function with given fields: timeSlots, date
func (_m *MockBConversionUtil) ConvertTimeSlotsStringToAppointments(timeSlots string, date *time.Time) *[]models.Appointment {
	ret := _m.Called(timeSlots, date)

	var r0 *[]models.Appointment
	if rf, ok := ret.Get(0).(func(string, *time.Time) *[]models.Appointment); ok {
		r0 = rf(timeSlots, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Appointment)
		}
	}

	return r0
}

// MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertTimeSlotsStringToAppointments'
type MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call struct {
	*mock.Call
}

// ConvertTimeSlotsStringToAppointments is a helper method to define mock.On call
//   - timeSlots string
//   - date *time.Time
func (_e *MockBConversionUtil_Expecter) ConvertTimeSlotsStringToAppointments(timeSlots interface{}, date interface{}) *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call {
	return &MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call{Call: _e.mock.On("ConvertTimeSlotsStringToAppointments", timeSlots, date)}
}

func (_c *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call) Run(run func(timeSlots string, date *time.Time)) *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*time.Time))
	})
	return _c
}

func (_c *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call) Return(_a0 *[]models.Appointment) *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call) RunAndReturn(run func(string, *time.Time) *[]models.Appointment) *MockBConversionUtil_ConvertTimeSlotsStringToAppointments_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBConversionUtil interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBConversionUtil creates a new instance of MockBConversionUtil. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBConversionUtil(t mockConstructorTestingTNewMockBConversionUtil) *MockBConversionUtil {
	mock := &MockBConversionUtil{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}