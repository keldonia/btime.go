package core

import (
	"fmt"
	"strings"

	"github.com/keldonia/btime.go/constants"
	"github.com/keldonia/btime.go/models"
	"github.com/keldonia/btime.go/utils"
)

//go:generate mockery --name BScheduler
type BScheduler interface {
	UpdateScheduleWithAppointmentSchedule(proposedAppointmentSchedule *models.AppointmentSchedule, schedule *models.Schedule) (*models.AppointmentSchedule, error)
	ConvertScheduleToAppointmentSchedule(schedule *models.Schedule) (*models.AppointmentSchedule, error)
	GetCurrentAvailability(schedule *models.Schedule) (*[]string, error)
	UpdateSchedule(proposedSchedule *models.Schedule, schedule *models.Schedule) (*models.Schedule, error)
	ProcessAppointments(appointments *[]models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error)
	ProcessAppointment(appointment *models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error)
	ComposeAppointments(appointment *models.Appointment) *models.AppointmentDuo
	HandleBookingUpdate(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error)
	HandleBookingUpdateBString(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error)
	DeleteAppointment(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error)
	DeleteAppointments(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error)
}

type BSchedulerImpl struct {
	bTimeFactory BTimeFactory
}

// Instantiates a new BScheduler, is responsible or maintaining of scheduling using binary Scheduler
//
// NB: A time interval of 5 is usually adequate for most implementations
//
// NB: The time interval must be a factor of 60,
//
//	ie. 1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, or 60
func NewBScheduler(timeInterval int) (BScheduler, error) {
	bTimeFactory, err := NewBTimeFactory(timeInterval)

	if err != nil {
		return nil, err
	}

	return &BSchedulerImpl{
		bTimeFactory: bTimeFactory,
	}, nil
}

// Tests a proposed appointment schedule update and updates the schedule,
// if theupdate is valid or throws an error if the update is not valid
func (bs *BSchedulerImpl) UpdateScheduleWithAppointmentSchedule(proposedAppointmentSchedule *models.AppointmentSchedule, schedule *models.Schedule) (*models.AppointmentSchedule, error) {
	scheduleAppointments := []models.Appointment{}

	for i := 0; i < len(*proposedAppointmentSchedule.Schedule); i++ {
		appointments := (*proposedAppointmentSchedule.Schedule)[i]
		scheduleAppointments = append(scheduleAppointments, appointments...)
	}

	proposedScheduleBStrings, err := bs.bTimeFactory.GenerateBStringFromAppointments(&scheduleAppointments)

	if err != nil {
		return nil, err
	}

	proposedSchedule := models.Schedule{
		Schedule:  proposedScheduleBStrings,
		Bookings:  schedule.Bookings,
		WeekStart: schedule.WeekStart,
	}

	updatedSchedule, err := bs.UpdateSchedule(&proposedSchedule, schedule)

	if err != nil {
		return nil, err
	}

	availability, err := bs.GetCurrentAvailability(schedule)

	if err != nil {
		return nil, err
	}

	appointmentSchedule := bs.bTimeFactory.ConvertScheduleToAppointmentSchedule(updatedSchedule, *availability)

	return appointmentSchedule, nil
}

// Takes a schedule and converts into an array of appointments for each date
//
// NB: This is a passthrough to the configured BTimeFactory
func (bs *BSchedulerImpl) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule) (*models.AppointmentSchedule, error) {
	availability, err := bs.GetCurrentAvailability(schedule)

	if err != nil {
		return nil, fmt.Errorf("BScheduler Error: Was unable to convert schedule to appointment schedule, as the bookings do not fit in the schedule")
	}

	appointmentSchedule := bs.bTimeFactory.ConvertScheduleToAppointmentSchedule(schedule, *availability)

	return appointmentSchedule, nil
}

// Takes a valid schedule and computes the remaining availability
// based on the total availability and current bookings, throws an error if an
// invalid scehdule is passed
func (bs *BSchedulerImpl) GetCurrentAvailability(schedule *models.Schedule) (*[]string, error) {
	totalRemainingAvailability := []string{}

	for i := 0; i < constants.DaysInWeek; i++ {
		availability := (*schedule.Schedule)[i]
		splitBookings := bs.bTimeFactory.TimeStringSplit((*schedule.Bookings)[i])
		splitAvailability := bs.bTimeFactory.TimeStringSplit(availability)
		calculatedAvailability := []string{}

		for j := 0; j < len(splitBookings); j++ {
			availabilityInterval, err := bs.bTimeFactory.ParseBString(splitAvailability[j])

			if err != nil {
				return nil, err
			}

			bBookingInterval, err := bs.bTimeFactory.ParseBString(splitBookings[j])

			if err != nil {
				return nil, err
			}

			// We bitwise NOT the availabilty interval here to test for outside overlap
			remainingAvailabilityMask, err := bs.bTimeFactory.TestViabilityAndCompute(^*availabilityInterval, *bBookingInterval)
			if err != nil {
				return nil, err
			}

			remainingAvailailiy := ^*remainingAvailabilityMask
			calculatedAvailability = append(calculatedAvailability, bs.bTimeFactory.DecimalToBString(float64(remainingAvailailiy)))
		}

		joinedRemainingAvailability := strings.Join(calculatedAvailability, "")
		totalRemainingAvailability = append(totalRemainingAvailability, joinedRemainingAvailability)
	}

	return &totalRemainingAvailability, nil
}

// Tests a propsoed schedule update and updates the schedule,
// if the update is valid or throws an error if the update is not valid
func (bs *BSchedulerImpl) UpdateSchedule(proposedSchedule *models.Schedule, schedule *models.Schedule) (*models.Schedule, error) {
	for i := 0; i < constants.DaysInWeek; i++ {
		// We test that no bookings fall outside of the scheduled availability
		proposed := (*proposedSchedule.Schedule)[i]
		splitBookings := bs.bTimeFactory.TimeStringSplit((*schedule.Bookings)[i])
		splitProposed := bs.bTimeFactory.TimeStringSplit(proposed)

		for j := 0; j < constants.HoursInDay; j++ {
			bBookingInterval, err := bs.bTimeFactory.ParseBString(splitBookings[j])

			if err != nil {
				return nil, err
			}

			proposedInterval, err := bs.bTimeFactory.ParseBString(splitProposed[j])
			if err != nil {
				return nil, err
			}

			// We bitwise NOT the proposed interval here to test for outside overlap
			_, err = bs.bTimeFactory.TestViabilityAndCompute(^*proposedInterval, *bBookingInterval)
			if err != nil {
				return nil, err
			}
		}

		joinedProposeded := strings.Join(splitProposed[0:constants.HoursInDay], "")
		joinedBookings := strings.Join(splitBookings[0:constants.HoursInDay], "")

		(*proposedSchedule.Schedule)[i] = joinedProposeded
		(*schedule.Bookings)[i] = joinedBookings
	}

	schedule.Schedule = proposedSchedule.Schedule

	return schedule, nil
}

// Takes an slice of appointments and update type and tests if the appointment updates are valid,
// if not it throws an error, if they are the schedule is updated
func (bs *BSchedulerImpl) ProcessAppointments(appointments *[]models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error) {
	appointmentsBStrings, err := bs.bTimeFactory.GenerateBStringFromAppointments(appointments)

	if err != nil {
		return nil, err
	}

	if actionType == constants.DELETE_APPOINTMENT {
		return bs.DeleteAppointments(*appointmentsBStrings, schedule)
	}

	if actionType == constants.BOOKING_UPDATE {
		return bs.HandleBookingUpdateBString(*appointmentsBStrings, schedule)
	}

	return nil, fmt.Errorf("BScheduler Error: Recieved invalid action type: %s", actionType)
}

// Takes an appointment and update type and tests if the appointment update is valid,
// if not it throws an error, if it is the schedule is updated
func (bs *BSchedulerImpl) ProcessAppointment(appointment *models.Appointment, schedule *models.Schedule, actionType constants.ScheduleAction) (*models.Schedule, error) {
	crossesDayBoundary := utils.CrossesDayBoundary(*appointment)
	var firstAppt *models.Appointment

	if crossesDayBoundary {
		appointmentDuo := bs.ComposeAppointments(appointment)

		firstAppt = appointmentDuo.SecondAppointment
		appointment = appointmentDuo.InitialAppointment
	}

	if actionType == constants.DELETE_APPOINTMENT {
		return bs.DeleteAppointment(appointment, schedule, firstAppt)
	}

	if actionType == constants.BOOKING_UPDATE {
		return bs.HandleBookingUpdate(appointment, schedule, firstAppt)
	}

	return nil, fmt.Errorf("BScheduler Error: Recieved invalid action type: %s", actionType)
}

// Utility function to split appointments that cross the day boundary
func (bs *BSchedulerImpl) ComposeAppointments(appointment *models.Appointment) *models.AppointmentDuo {
	utcAppt := utils.EnforceUTC(appointment)
	utcStartTime := utcAppt.StartTime
	utcEndTime := utcAppt.EndTime

	// Clone Appt
	initialAppointment := models.Appointment{
		StartTime: utcStartTime,
		EndTime:   utils.GetUTCDateEnd(utcStartTime, 59),
	}

	SecondAppointment := models.Appointment{
		StartTime: utils.GetUTCDateStart(*utcEndTime),
		EndTime:   utcEndTime,
	}

	return &models.AppointmentDuo{
		InitialAppointment: &initialAppointment,
		SecondAppointment:  &SecondAppointment,
	}
}

// Takes an appointment and tests if the appointment update is valid, if not it throws an error, if it is the schedule is updated
func (bs *BSchedulerImpl) HandleBookingUpdate(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error) {
	var startDay int = int(appointment.StartTime.UTC().Weekday())
	endDay := int(appointment.EndTime.UTC().Weekday())

	if firstAppt != nil {
		startDay = int(firstAppt.StartTime.UTC().Weekday())
		firstAppointmentBString, err := bs.bTimeFactory.GenerateBString(firstAppt)

		if err != nil {
			return nil, err
		}

		tempBookings, err := bs.bTimeFactory.ModifyScheduleAndBooking(
			(*schedule.Bookings)[startDay],
			(*schedule.Schedule)[startDay],
			*firstAppointmentBString,
		)

		if err != nil {
			return nil, err
		}

		(*schedule.Bookings)[startDay] = *tempBookings
	}

	apptBString, err := bs.bTimeFactory.GenerateBString(appointment)

	if err != nil {
		return nil, err
	}

	tempBookings, err := bs.bTimeFactory.ModifyScheduleAndBooking(
		(*schedule.Bookings)[endDay],
		(*schedule.Schedule)[endDay],
		*apptBString,
	)

	if err != nil {
		return nil, err
	}

	(*schedule.Bookings)[endDay] = *tempBookings

	return schedule, nil
}

// Takes an array of appointments and tests if the appointment update are valid, if not it throws an error, if they are the schedule is updated
func (bs *BSchedulerImpl) HandleBookingUpdateBString(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error) {
	bookings := []string{}

	for i := 0; i < constants.DaysInWeek; i++ {
		tempBookings, err := bs.bTimeFactory.ModifyScheduleAndBooking(
			(*schedule.Bookings)[i],
			(*schedule.Schedule)[i],
			appointmentsBStrings[i],
		)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, *tempBookings)
	}

	*schedule.Bookings = bookings

	return schedule, nil
}

// Takes an appointment and tests if the appointment to delete is valid, if not throws an error,
// if it is the schedule is updated to reflect the deletion
func (bs *BSchedulerImpl) DeleteAppointment(appointment *models.Appointment, schedule *models.Schedule, firstAppt *models.Appointment) (*models.Schedule, error) {
	var startDay int = bs.bTimeFactory.FindWeekDay(appointment.StartTime)
	endDay := bs.bTimeFactory.FindWeekDay(appointment.EndTime)

	if firstAppt != nil {
		startDay = bs.bTimeFactory.FindWeekDay(firstAppt.StartTime)
		firstAppointmentCalculated, err := bs.bTimeFactory.DeleteAppointment(firstAppt, (*schedule.Bookings)[startDay])

		if err != nil {
			return nil, err
		}

		(*schedule.Bookings)[startDay] = *firstAppointmentCalculated
	}

	mainCalculated, err := bs.bTimeFactory.DeleteAppointment(appointment, (*&*schedule.Bookings)[endDay])

	if err != nil {
		return nil, err
	}

	(*schedule.Bookings)[endDay] = *mainCalculated

	return schedule, nil
}

// Takes an array of appointments and tests if the appointments to delete are valid, if not throws an error,
// if they are the schedule is updated to reflect the deletion
func (bs *BSchedulerImpl) DeleteAppointments(appointmentsBStrings []string, schedule *models.Schedule) (*models.Schedule, error) {
	bookings := []string{}

	for i := 0; i < constants.DaysInWeek; i++ {
		calculatedSchedule, err := bs.bTimeFactory.DeleteAppointmentBString(appointmentsBStrings[i], (*schedule.Bookings)[i])

		if err != nil {
			return nil, fmt.Errorf("BSchedule Error: interval to delete occurs outside of schedule on day %d of the week starting on %s", i, schedule.WeekStart.UTC().GoString())
		}

		bookings = append(bookings, *calculatedSchedule)
	}

	schedule.Bookings = &bookings

	return schedule, nil
}
