// Code generated by mockery v2.20.2. DO NOT EDIT.

package core

import (
	constants "github.com/keldonia/btime.go/constants"
	mock "github.com/stretchr/testify/mock"

	models "github.com/keldonia/btime.go/models"
)

// MockBScheduler is an autogenerated mock type for the BScheduler type
type MockBScheduler struct {
	mock.Mock
}

type MockBScheduler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBScheduler) EXPECT() *MockBScheduler_Expecter {
	return &MockBScheduler_Expecter{mock: &_m.Mock}
}

// ComposeAppointments provides a mock function with given fields: appointment
func (_m *MockBScheduler) ComposeAppointments(appointment *models.Appointment) *models.AppointmentDuo {
	ret := _m.Called(appointment)

	var r0 *models.AppointmentDuo
	if rf, ok := ret.Get(0).(func(*models.Appointment) *models.AppointmentDuo); ok {
		r0 = rf(appointment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AppointmentDuo)
		}
	}

	return r0
}

// MockBScheduler_ComposeAppointments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComposeAppointments'
type MockBScheduler_ComposeAppointments_Call struct {
	*mock.Call
}

// ComposeAppointments is a helper method to define mock.On call
//   - appointment *models.Appointment
func (_e *MockBScheduler_Expecter) ComposeAppointments(appointment interface{}) *MockBScheduler_ComposeAppointments_Call {
	return &MockBScheduler_ComposeAppointments_Call{Call: _e.mock.On("ComposeAppointments", appointment)}
}

func (_c *MockBScheduler_ComposeAppointments_Call) Run(run func(appointment *models.Appointment)) *MockBScheduler_ComposeAppointments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment))
	})
	return _c
}

func (_c *MockBScheduler_ComposeAppointments_Call) Return(_a0 *models.AppointmentDuo) *MockBScheduler_ComposeAppointments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBScheduler_ComposeAppointments_Call) RunAndReturn(run func(*models.Appointment) *models.AppointmentDuo) *MockBScheduler_ComposeAppointments_Call {
	_c.Call.Return(run)
	return _c
}

// ConvertScheduleToAppointmentSchedule provides a mock function with given fields: schedule
func (_m *MockBScheduler) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule) (*models.AppointmentSchedule, error) {
	ret := _m.Called(schedule)

	var r0 *models.AppointmentSchedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Schedule) (*models.AppointmentSchedule, error)); ok {
		return rf(schedule)
	}
	if rf, ok := ret.Get(0).(func(*models.Schedule) *models.AppointmentSchedule); ok {
		r0 = rf(schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AppointmentSchedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Schedule) error); ok {
		r1 = rf(schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_ConvertScheduleToAppointmentSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertScheduleToAppointmentSchedule'
type MockBScheduler_ConvertScheduleToAppointmentSchedule_Call struct {
	*mock.Call
}

// ConvertScheduleToAppointmentSchedule is a helper method to define mock.On call
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) ConvertScheduleToAppointmentSchedule(schedule interface{}) *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call {
	return &MockBScheduler_ConvertScheduleToAppointmentSchedule_Call{Call: _e.mock.On("ConvertScheduleToAppointmentSchedule", schedule)}
}

func (_c *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call) Run(run func(schedule *models.Schedule)) *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call) Return(_a0 *models.AppointmentSchedule, _a1 error) *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call) RunAndReturn(run func(*models.Schedule) (*models.AppointmentSchedule, error)) *MockBScheduler_ConvertScheduleToAppointmentSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAppointment provides a mock function with given fields: appointment, schedule, firstAppt
func (_m *MockBScheduler) DeleteAppointment(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error) {
	ret := _m.Called(appointment, schedule, firstAppt)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, *models.Appointment) (*models.Schedule, error)); ok {
		return rf(appointment, schedule, firstAppt)
	}
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, *models.Appointment) *models.Schedule); ok {
		r0 = rf(appointment, schedule, firstAppt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Appointment, *models.Schedule, *models.Appointment) error); ok {
		r1 = rf(appointment, schedule, firstAppt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_DeleteAppointment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAppointment'
type MockBScheduler_DeleteAppointment_Call struct {
	*mock.Call
}

// DeleteAppointment is a helper method to define mock.On call
//   - appointment *models.Appointment
//   - schedule *models.Schedule
//   - firstAppt *models.Appointment
func (_e *MockBScheduler_Expecter) DeleteAppointment(appointment interface{}, schedule interface{}, firstAppt interface{}) *MockBScheduler_DeleteAppointment_Call {
	return &MockBScheduler_DeleteAppointment_Call{Call: _e.mock.On("DeleteAppointment", appointment, schedule, firstAppt)}
}

func (_c *MockBScheduler_DeleteAppointment_Call) Run(run func(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment)) *MockBScheduler_DeleteAppointment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment), args[1].(*models.Schedule), args[2].(*models.Appointment))
	})
	return _c
}

func (_c *MockBScheduler_DeleteAppointment_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_DeleteAppointment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_DeleteAppointment_Call) RunAndReturn(run func(*models.Appointment, *models.Schedule, *models.Appointment) (*models.Schedule, error)) *MockBScheduler_DeleteAppointment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAppointments provides a mock function with given fields: appointmentsBStrings, schedule
func (_m *MockBScheduler) DeleteAppointments(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error) {
	ret := _m.Called(appointmentsBStrings, schedule)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, *models.Schedule) (*models.Schedule, error)); ok {
		return rf(appointmentsBStrings, schedule)
	}
	if rf, ok := ret.Get(0).(func([]string, *models.Schedule) *models.Schedule); ok {
		r0 = rf(appointmentsBStrings, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, *models.Schedule) error); ok {
		r1 = rf(appointmentsBStrings, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_DeleteAppointments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAppointments'
type MockBScheduler_DeleteAppointments_Call struct {
	*mock.Call
}

// DeleteAppointments is a helper method to define mock.On call
//   - appointmentsBStrings []string
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) DeleteAppointments(appointmentsBStrings interface{}, schedule interface{}) *MockBScheduler_DeleteAppointments_Call {
	return &MockBScheduler_DeleteAppointments_Call{Call: _e.mock.On("DeleteAppointments", appointmentsBStrings, schedule)}
}

func (_c *MockBScheduler_DeleteAppointments_Call) Run(run func(appointmentsBStrings []string, schedule *models.Schedule)) *MockBScheduler_DeleteAppointments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_DeleteAppointments_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_DeleteAppointments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_DeleteAppointments_Call) RunAndReturn(run func([]string, *models.Schedule) (*models.Schedule, error)) *MockBScheduler_DeleteAppointments_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrentAvailability provides a mock function with given fields: schedule
func (_m *MockBScheduler) GetCurrentAvailability(schedule *models.Schedule) (*[]string, error) {
	ret := _m.Called(schedule)

	var r0 *[]string
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Schedule) (*[]string, error)); ok {
		return rf(schedule)
	}
	if rf, ok := ret.Get(0).(func(*models.Schedule) *[]string); ok {
		r0 = rf(schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Schedule) error); ok {
		r1 = rf(schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_GetCurrentAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentAvailability'
type MockBScheduler_GetCurrentAvailability_Call struct {
	*mock.Call
}

// GetCurrentAvailability is a helper method to define mock.On call
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) GetCurrentAvailability(schedule interface{}) *MockBScheduler_GetCurrentAvailability_Call {
	return &MockBScheduler_GetCurrentAvailability_Call{Call: _e.mock.On("GetCurrentAvailability", schedule)}
}

func (_c *MockBScheduler_GetCurrentAvailability_Call) Run(run func(schedule *models.Schedule)) *MockBScheduler_GetCurrentAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_GetCurrentAvailability_Call) Return(_a0 *[]string, _a1 error) *MockBScheduler_GetCurrentAvailability_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_GetCurrentAvailability_Call) RunAndReturn(run func(*models.Schedule) (*[]string, error)) *MockBScheduler_GetCurrentAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// HandleBookingUpdate provides a mock function with given fields: appointment, schedule, firstAppt
func (_m *MockBScheduler) HandleBookingUpdate(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error) {
	ret := _m.Called(appointment, schedule, firstAppt)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, *models.Appointment) (*models.Schedule, error)); ok {
		return rf(appointment, schedule, firstAppt)
	}
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, *models.Appointment) *models.Schedule); ok {
		r0 = rf(appointment, schedule, firstAppt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Appointment, *models.Schedule, *models.Appointment) error); ok {
		r1 = rf(appointment, schedule, firstAppt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_HandleBookingUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleBookingUpdate'
type MockBScheduler_HandleBookingUpdate_Call struct {
	*mock.Call
}

// HandleBookingUpdate is a helper method to define mock.On call
//   - appointment *models.Appointment
//   - schedule *models.Schedule
//   - firstAppt *models.Appointment
func (_e *MockBScheduler_Expecter) HandleBookingUpdate(appointment interface{}, schedule interface{}, firstAppt interface{}) *MockBScheduler_HandleBookingUpdate_Call {
	return &MockBScheduler_HandleBookingUpdate_Call{Call: _e.mock.On("HandleBookingUpdate", appointment, schedule, firstAppt)}
}

func (_c *MockBScheduler_HandleBookingUpdate_Call) Run(run func(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment)) *MockBScheduler_HandleBookingUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment), args[1].(*models.Schedule), args[2].(*models.Appointment))
	})
	return _c
}

func (_c *MockBScheduler_HandleBookingUpdate_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_HandleBookingUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_HandleBookingUpdate_Call) RunAndReturn(run func(*models.Appointment, *models.Schedule, *models.Appointment) (*models.Schedule, error)) *MockBScheduler_HandleBookingUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// HandleBookingUpdateBString provides a mock function with given fields: appointmentsBStrings, schedule
func (_m *MockBScheduler) HandleBookingUpdateBString(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error) {
	ret := _m.Called(appointmentsBStrings, schedule)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, *models.Schedule) (*models.Schedule, error)); ok {
		return rf(appointmentsBStrings, schedule)
	}
	if rf, ok := ret.Get(0).(func([]string, *models.Schedule) *models.Schedule); ok {
		r0 = rf(appointmentsBStrings, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, *models.Schedule) error); ok {
		r1 = rf(appointmentsBStrings, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_HandleBookingUpdateBString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleBookingUpdateBString'
type MockBScheduler_HandleBookingUpdateBString_Call struct {
	*mock.Call
}

// HandleBookingUpdateBString is a helper method to define mock.On call
//   - appointmentsBStrings []string
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) HandleBookingUpdateBString(appointmentsBStrings interface{}, schedule interface{}) *MockBScheduler_HandleBookingUpdateBString_Call {
	return &MockBScheduler_HandleBookingUpdateBString_Call{Call: _e.mock.On("HandleBookingUpdateBString", appointmentsBStrings, schedule)}
}

func (_c *MockBScheduler_HandleBookingUpdateBString_Call) Run(run func(appointmentsBStrings []string, schedule *models.Schedule)) *MockBScheduler_HandleBookingUpdateBString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string), args[1].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_HandleBookingUpdateBString_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_HandleBookingUpdateBString_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_HandleBookingUpdateBString_Call) RunAndReturn(run func([]string, *models.Schedule) (*models.Schedule, error)) *MockBScheduler_HandleBookingUpdateBString_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessAppointment provides a mock function with given fields: appointment, schedule, actionType
func (_m *MockBScheduler) ProcessAppointment(appointment *models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error) {
	ret := _m.Called(appointment, schedule, actionType)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, constants.ScheduleAction) (*models.Schedule, error)); ok {
		return rf(appointment, schedule, actionType)
	}
	if rf, ok := ret.Get(0).(func(*models.Appointment, *models.Schedule, constants.ScheduleAction) *models.Schedule); ok {
		r0 = rf(appointment, schedule, actionType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Appointment, *models.Schedule, constants.ScheduleAction) error); ok {
		r1 = rf(appointment, schedule, actionType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_ProcessAppointment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessAppointment'
type MockBScheduler_ProcessAppointment_Call struct {
	*mock.Call
}

// ProcessAppointment is a helper method to define mock.On call
//   - appointment *models.Appointment
//   - schedule *models.Schedule
//   - actionType constants.ScheduleAction
func (_e *MockBScheduler_Expecter) ProcessAppointment(appointment interface{}, schedule interface{}, actionType interface{}) *MockBScheduler_ProcessAppointment_Call {
	return &MockBScheduler_ProcessAppointment_Call{Call: _e.mock.On("ProcessAppointment", appointment, schedule, actionType)}
}

func (_c *MockBScheduler_ProcessAppointment_Call) Run(run func(appointment *models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction)) *MockBScheduler_ProcessAppointment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Appointment), args[1].(*models.Schedule), args[2].(constants.ScheduleAction))
	})
	return _c
}

func (_c *MockBScheduler_ProcessAppointment_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_ProcessAppointment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_ProcessAppointment_Call) RunAndReturn(run func(*models.Appointment, *models.Schedule, constants.ScheduleAction) (*models.Schedule, error)) *MockBScheduler_ProcessAppointment_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessAppointments provides a mock function with given fields: appointments, schedule, actionType
func (_m *MockBScheduler) ProcessAppointments(appointments *[]models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error) {
	ret := _m.Called(appointments, schedule, actionType)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*[]models.Appointment, *models.Schedule, constants.ScheduleAction) (*models.Schedule, error)); ok {
		return rf(appointments, schedule, actionType)
	}
	if rf, ok := ret.Get(0).(func(*[]models.Appointment, *models.Schedule, constants.ScheduleAction) *models.Schedule); ok {
		r0 = rf(appointments, schedule, actionType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*[]models.Appointment, *models.Schedule, constants.ScheduleAction) error); ok {
		r1 = rf(appointments, schedule, actionType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_ProcessAppointments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessAppointments'
type MockBScheduler_ProcessAppointments_Call struct {
	*mock.Call
}

// ProcessAppointments is a helper method to define mock.On call
//   - appointments *[]models.Appointment
//   - schedule *models.Schedule
//   - actionType constants.ScheduleAction
func (_e *MockBScheduler_Expecter) ProcessAppointments(appointments interface{}, schedule interface{}, actionType interface{}) *MockBScheduler_ProcessAppointments_Call {
	return &MockBScheduler_ProcessAppointments_Call{Call: _e.mock.On("ProcessAppointments", appointments, schedule, actionType)}
}

func (_c *MockBScheduler_ProcessAppointments_Call) Run(run func(appointments *[]models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction)) *MockBScheduler_ProcessAppointments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*[]models.Appointment), args[1].(*models.Schedule), args[2].(constants.ScheduleAction))
	})
	return _c
}

func (_c *MockBScheduler_ProcessAppointments_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_ProcessAppointments_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_ProcessAppointments_Call) RunAndReturn(run func(*[]models.Appointment, *models.Schedule, constants.ScheduleAction) (*models.Schedule, error)) *MockBScheduler_ProcessAppointments_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSchedule provides a mock function with given fields: proposedSchedule, schedule
func (_m *MockBScheduler) UpdateSchedule(proposedSchedule *models.Schedule, schedule *models.Schedule) (*models.Schedule, error) {
	ret := _m.Called(proposedSchedule, schedule)

	var r0 *models.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Schedule, *models.Schedule) (*models.Schedule, error)); ok {
		return rf(proposedSchedule, schedule)
	}
	if rf, ok := ret.Get(0).(func(*models.Schedule, *models.Schedule) *models.Schedule); ok {
		r0 = rf(proposedSchedule, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Schedule, *models.Schedule) error); ok {
		r1 = rf(proposedSchedule, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_UpdateSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSchedule'
type MockBScheduler_UpdateSchedule_Call struct {
	*mock.Call
}

// UpdateSchedule is a helper method to define mock.On call
//   - proposedSchedule *models.Schedule
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) UpdateSchedule(proposedSchedule interface{}, schedule interface{}) *MockBScheduler_UpdateSchedule_Call {
	return &MockBScheduler_UpdateSchedule_Call{Call: _e.mock.On("UpdateSchedule", proposedSchedule, schedule)}
}

func (_c *MockBScheduler_UpdateSchedule_Call) Run(run func(proposedSchedule *models.Schedule, schedule *models.Schedule)) *MockBScheduler_UpdateSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Schedule), args[1].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_UpdateSchedule_Call) Return(_a0 *models.Schedule, _a1 error) *MockBScheduler_UpdateSchedule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_UpdateSchedule_Call) RunAndReturn(run func(*models.Schedule, *models.Schedule) (*models.Schedule, error)) *MockBScheduler_UpdateSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateScheduleWithAppointmentSchedule provides a mock function with given fields: proposedAppointmentSchedule, schedule
func (_m *MockBScheduler) UpdateScheduleWithAppointmentSchedule(proposedAppointmentSchedule *models.AppointmentSchedule, schedule *models.Schedule) (*models.AppointmentSchedule, error) {
	ret := _m.Called(proposedAppointmentSchedule, schedule)

	var r0 *models.AppointmentSchedule
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.AppointmentSchedule, *models.Schedule) (*models.AppointmentSchedule, error)); ok {
		return rf(proposedAppointmentSchedule, schedule)
	}
	if rf, ok := ret.Get(0).(func(*models.AppointmentSchedule, *models.Schedule) *models.AppointmentSchedule); ok {
		r0 = rf(proposedAppointmentSchedule, schedule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AppointmentSchedule)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.AppointmentSchedule, *models.Schedule) error); ok {
		r1 = rf(proposedAppointmentSchedule, schedule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateScheduleWithAppointmentSchedule'
type MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call struct {
	*mock.Call
}

// UpdateScheduleWithAppointmentSchedule is a helper method to define mock.On call
//   - proposedAppointmentSchedule *models.AppointmentSchedule
//   - schedule *models.Schedule
func (_e *MockBScheduler_Expecter) UpdateScheduleWithAppointmentSchedule(proposedAppointmentSchedule interface{}, schedule interface{}) *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call {
	return &MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call{Call: _e.mock.On("UpdateScheduleWithAppointmentSchedule", proposedAppointmentSchedule, schedule)}
}

func (_c *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call) Run(run func(proposedAppointmentSchedule *models.AppointmentSchedule, schedule *models.Schedule)) *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.AppointmentSchedule), args[1].(*models.Schedule))
	})
	return _c
}

func (_c *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call) Return(_a0 *models.AppointmentSchedule, _a1 error) *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call) RunAndReturn(run func(*models.AppointmentSchedule, *models.Schedule) (*models.AppointmentSchedule, error)) *MockBScheduler_UpdateScheduleWithAppointmentSchedule_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBScheduler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBScheduler creates a new instance of MockBScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBScheduler(t mockConstructorTestingTNewMockBScheduler) *MockBScheduler {
	mock := &MockBScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
