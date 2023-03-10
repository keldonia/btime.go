package core

import (
	"time"

	"github.com/keldonia/btime.go/models"
)

//go:generate mockery --name BTimeFactory
type BTimeFactory interface {
	ParseBString(bString string) (*int64, error)
	GenerateBString(appt *models.Appointment) (*string, error)
	GenerateBStringFromAppointments(appointments *[]models.Appointment) (*[]string, error)
	TimeStringSplit(scheduleString string) []string
	DecimalToBString(decimal float64) string
	TestViabilityAndCompute(binary1 int64, binary2 int64) (*int64, error)
	DeleteAppointment(timeSlotToDelete *models.Appointment, scheduleSlot string) (*string, error)
	DeleteAppointmentBString(timeSlotToDelete string, scheduleSlot string) (*string, error)
	ModifyScheduleAndBooking(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error)
	ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule
	FindWeekDay(date *time.Time) int
}

type BTimeFactoryImpl struct {
	bTimeConfig        *BTimeConfig
	bStringUtil        BStringUtil
	bScheduleUtil      BScheduleUtil
	bConversionUtil    BConversionUtil
	bPointerCalculator BPointerCalculator
}

// Instantiates a new BTimeFactory, which manages and exposes various binary scheduling and string utils
//
// NB: Typically a temporal resolution of 5 mins is sufficient,
// as it constitutes the smallest billable unit in most juristictions
//
// NB: The time interval must be a factor of 60,
//
//	ie. 1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, or 60
func NewBTimeFactory(timeInterval int) (BTimeFactory, error) {
	bTimeConfig, err := BuildConfigFromTimeInterval(timeInterval)

	if err != nil {
		return nil, err
	}

	// We already checked that bTimeConfig will not be nil
	bStringUtil, _ := NewBStringUtil(bTimeConfig)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)
	bConversionUtil, _ := NewBConversionUtil(bTimeConfig)
	bPointerCalculator, _ := NewBPointerCalculator(bTimeConfig)

	return &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}, nil
}

// Converts bString representation of a number into a number for calculation purposes
//
// NB: This is a passthrough to the configured bStringUtil
func (btf *BTimeFactoryImpl) ParseBString(bString string) (*int64, error) {
	num, err := btf.bStringUtil.ParseBString(bString)

	if err != nil {
		return nil, err
	}

	return num, nil
}

// Generates a bString representation of a given appointment, assuming it is valid.
// If the appointment is invalid, it throws an error
//
// NB: This is a passthrough to the configured bStringUtil
func (btf *BTimeFactoryImpl) GenerateBString(appt *models.Appointment) (*string, error) {
	bString, err := btf.bStringUtil.GenerateBString(appt)

	if err != nil {
		return nil, err
	}

	return bString, nil
}

// Generates a bString representation of a given array of appointments, assuming it is valid.
// If the appointment is invalid, it will throw an error
//
// NB: This method generates a representation of the entire week
//
// NB: Assumes appointments in array don't overlap
//
// NB: This is a passthrough to the configured bStringUtil
func (btf *BTimeFactoryImpl) GenerateBStringFromAppointments(appointments *[]models.Appointment) (*[]string, error) {
	bString, err := btf.bStringUtil.GenerateBStringFromAppointments(appointments)

	if err != nil {
		return nil, err
	}

	return bString, nil
}

// Splits each schedule bString into a string of length defined in the regex
//
// NB: This is a passthrough to the configured bStringUtil
func (btf *BTimeFactoryImpl) TimeStringSplit(scheduleString string) []string {
	return btf.bStringUtil.TimeStringSplit(scheduleString)
}

// Converts number into a bString representation with the given scheduling interval
//
// NB: This is a passthrough to the configured bStringUtil
func (btf *BTimeFactoryImpl) DecimalToBString(decimal float64) string {
	return btf.bStringUtil.DecimalToBString(decimal)
}

// Tests that two time intervals do not overlap, throwing an error if they do
//
// NB: This is a passthrough to the configured bScheduleUtil
func (btf *BTimeFactoryImpl) TestViabilityAndCompute(binary1 int64, binary2 int64) (*int64, error) {
	computed, err := btf.bScheduleUtil.TestViabilityAndCompute(binary1, binary2)

	if err != nil {
		return nil, err
	}

	return computed, nil
}

// Tests removal a give time slot from a given time interval and if valid removes it, else throws an error
//
// NB: This is also used for calculating remaining availability
//
// NB: This is a passthrough to the configured bScheduleUtil
func (btf *BTimeFactoryImpl) DeleteAppointment(timeSlotToDelete *models.Appointment, scheduleSlot string) (*string, error) {
	updatedApptBString, err := btf.bScheduleUtil.DeleteAppointment(timeSlotToDelete, scheduleSlot)

	if err != nil {
		return nil, err
	}

	return updatedApptBString, nil
}

// Tests removal a give time slot from a given time interval and if valid removes it, else throws an error
//
// NB: This is also used for calculating remaining availability
//
// NB: This is a passthrough to the configured bScheduleUtil
func (btf *BTimeFactoryImpl) DeleteAppointmentBString(timeSlotToDelete string, scheduleSlot string) (*string, error) {
	updatedApptBString, err := btf.bScheduleUtil.DeleteAppointmentBString(timeSlotToDelete, scheduleSlot)

	if err != nil {
		return nil, err
	}

	return updatedApptBString, nil
}

// Tests that an timeSlot does not overlap with another timeSlot, if it does not overlap, the timeSlot is added to the bookings,
// else throw an error.  Additionally, this method checks that the timeslot is within availabilities (test)
//
// NB: If testing a booking update, test that booking fits in avail
// This means that bookingsUpdate the inputs are (bookings, bookings, appt)
//
// NB: This is a passthrough to the configured bScheduleUtil
func (btf *BTimeFactoryImpl) ModifyScheduleAndBooking(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	modifiedBookings, err := btf.bScheduleUtil.ModifyScheduleAndBooking(scheduleBStringToModify, scheduleBStringToTest, appt)

	if err != nil {
		return nil, err
	}

	return *&modifiedBookings, err
}

// Takes a schedule and availabilty converting them into an array of appointments for each date
//
// NB: This is a passthrough to the configured BConversionUtil
func (btf *BTimeFactoryImpl) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule {
	return btf.bConversionUtil.ConvertScheduleToAppointmentSchedule(schedule, availability)
}

// Finds the numeric day of the week with 0 = Sunday
func (btf *BTimeFactoryImpl) FindWeekDay(date *time.Time) int {
	return btf.bPointerCalculator.FindWeekDay(date)
}
