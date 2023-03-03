package core

import (
	"github.com/keldonia/btime.go/models"
)

type BTimeFactory struct {
	bTimeConfig     *BTimeConfig
	bStringUtil     *BStringUtil
	bScheduleUtil   *BScheduleUtil
	bConversionUtil *BConversionUtil
}

func NewBTimeFactory(timeInterval int) (*BTimeFactory, error) {
	bTimeConfig, err := BuildConfigFromTimeInterval(timeInterval)

	if err != nil {
		return nil, err
	}

	bStringUtil, err := NewBStringUtil(bTimeConfig)

	if err != nil {
		return nil, err
	}

	bScheduleUtil, err := NewBScheduleUtil(bTimeConfig)

	if err != nil {
		return nil, err
	}

	bConversionUtil, err := NewBConversionUtil(bTimeConfig)

	if err != nil {
		return nil, err
	}

	return &BTimeFactory{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}, nil
}

func (btf *BTimeFactory) ParseBString(bString string) (*int64, error) {
	num, err := btf.bStringUtil.ParseBString(bString)

	if err != nil {
		return nil, err
	}

	return num, nil
}

func (btf *BTimeFactory) GenerateBString(appt *models.Appointment) (*string, error) {
	bString, err := btf.bStringUtil.GenerateBString(appt)

	if err != nil {
		return nil, err
	}

	return bString, nil
}

func (btf *BTimeFactory) GenerateBStringFromAppointments(appointments *[]models.Appointment) (*[]string, error) {
	bString, err := btf.bStringUtil.GenerateBStringFromAppointments(appointments)

	if err != nil {
		return nil, err
	}

	return bString, nil
}

func (btf *BTimeFactory) TimeStringSplit(scheduleString string) []string {
	return btf.bStringUtil.TimeStringSplit(scheduleString)
}

func (btf *BTimeFactory) DecimalToBString(decimal float64) string {
	return btf.bStringUtil.DecimalToBString(decimal)
}

func (btf *BTimeFactory) TestViabilityAndCompute(binary1 int64, binary2 int64) (*int64, error) {
	computed, err := btf.bScheduleUtil.TestViabilityAndCompute(binary1, binary2)

	if err != nil {
		return nil, err
	}

	return computed, nil
}

func (btf *BTimeFactory) DeleteAppointment(timeSlotToDelete *models.Appointment, scheduleSlot string) (*string, error) {
	updatedApptBString, err := btf.bScheduleUtil.DeleteAppointment(timeSlotToDelete, scheduleSlot)

	if err != nil {
		return nil, err
	}

	return updatedApptBString, nil
}

func (btf *BTimeFactory) DeleteAppointmentBString(timeSlotToDelete string, scheduleSlot string) (*string, error) {
	updatedApptBString, err := btf.bScheduleUtil.DeleteAppointmentBString(timeSlotToDelete, scheduleSlot)

	if err != nil {
		return nil, err
	}

	return updatedApptBString, nil
}

func (btf *BTimeFactory) ModifyScheduleAndBooking(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	modifiedBookings, err := btf.bScheduleUtil.ModifyScheduleAndBooking(scheduleBStringToModify, scheduleBStringToTest, appt)

	if err != nil {
		return nil, err
	}

	return *&modifiedBookings, err
}

func (btf *BTimeFactory) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule {
	return btf.bConversionUtil.ConvertScheduleToAppointmentSchedule(schedule, availability)
}
