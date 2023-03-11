package core

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/keldonia/btime.go/constants"
	"github.com/keldonia/btime.go/models"
)

func generateTimeSet(timeInterval int, appointments []models.Appointment) []string {
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	timeSet := []string{}
	for i := 0; i < len(appointments); i++ {
		appt := appointments[i]
		apptStr, _ := bStringUtil.GenerateBString(&appt)
		timeSet = append(timeSet, *apptStr)
	}

	for i := 0; i < 7-len(timeSet); i++ {
		timeSet = append(timeSet, bTimeConfig.EmptyDay)
	}

	return timeSet
}

func generateSchedule(schedule []string, bookings []string, weekStart *time.Time) *models.Schedule {
	return &models.Schedule{
		Schedule:  &schedule,
		Bookings:  &bookings,
		WeekStart: weekStart,
	}
}

func generateEmptyWeekFromConfig(bTimeConfig BTimeConfig) *[]string {
	return &[]string{
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
		bTimeConfig.EmptyDay,
	}
}

func TestNewBSchedulerInvalid(t *testing.T) {
	timeInterval := -1

	bScheduler, err := NewBScheduler(timeInterval)
	expectedErr := fmt.Sprintf("[BConfig] received an invalid time interval: %d", timeInterval)

	if bScheduler != nil {
		t.Fatalf("expected bScheduler to be nil")
	}

	if err.Error() != expectedErr {
		t.Fatalf("expected error: %s, received: %s", expectedErr, err.Error())
	}
}

func TestNewBSchedulerValidInput(t *testing.T) {
	timeInterval := 5

	bScheduler, err := NewBScheduler(timeInterval)

	if err != nil {
		t.Fatalf("expected no error, received: %s", err.Error())
	}

	if bScheduler == nil {
		t.Fatalf("expected bScheduler to be generated")
	}
}

func TestUpdateSchedule(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	emptyWeek := generateEmptyWeekFromConfig(*bTimeConfig)
	emptyDay := bTimeConfig.EmptyDay
	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-02T00:00:00Z")

	type test struct {
		Name                  string
		ScheduledAvailability []models.Appointment
		Bookings              []models.Appointment
		ProposedSchedule      []models.Appointment
		Error                 bool
		ExpectedError         string
		ExtendWeek            bool
		ExtendDay             bool
		InvalidBookings       bool
		InvalidProposed       bool
	}

	tests := []test{
		{
			Name: "should return the modified schedule if the current bookings are contained within the proposed availability",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           false,
			ExpectedError:   "",
			ExtendWeek:      false,
			ExtendDay:       false,
			InvalidBookings: false,
			InvalidProposed: false,
		},
		{
			Name: "should return the modified schedule if the current bookings are contained within the proposed availability, ignoring if the schedule has additional days",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           false,
			ExpectedError:   "",
			ExtendWeek:      true,
			ExtendDay:       false,
			InvalidBookings: false,
			InvalidProposed: false,
		},
		{
			Name: "should return an error if the current bookings are not contained within the proposed availability, empty hour",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(12, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           true,
			ExpectedError:   "BScheduleUtil Error: Time intervals overlap.",
			ExtendWeek:      false,
			ExtendDay:       false,
			InvalidBookings: false,
			InvalidProposed: false,
		},
		{
			Name: "should return an error if the current bookings are not contained within the proposed availability, non-empty hour",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(12, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(12, 30, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           true,
			ExpectedError:   "BScheduleUtil Error: Time intervals overlap.",
			ExtendWeek:      false,
			ExtendDay:       false,
			InvalidBookings: false,
			InvalidProposed: false,
		},
		{
			Name: "should ignore any additional intervals outside of the day",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:         false,
			ExpectedError: "",
			ExtendWeek:    false,
			ExtendDay:     true,
		},
		{
			Name: "should throw if bookings are invalid",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           true,
			ExpectedError:   "strconv.ParseInt: parsing \"00000w\": invalid syntax",
			ExtendWeek:      false,
			ExtendDay:       false,
			InvalidBookings: true,
			InvalidProposed: false,
		},
		{
			Name: "should throw if bookings are invalid",
			ScheduledAvailability: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 17, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 17, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 17, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 17, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 17, 0, 8, 8),
			},
			Bookings: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3),
			},
			ProposedSchedule: []models.Appointment{
				generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2),
				generateApptFromHoursAndMins(9, 0, 18, 0, 3, 3),
				generateApptFromHoursAndMins(9, 0, 18, 0, 4, 4),
				generateApptFromHoursAndMins(9, 0, 18, 0, 5, 5),
				generateApptFromHoursAndMins(9, 0, 18, 0, 6, 6),
				generateApptFromHoursAndMins(9, 0, 18, 0, 7, 7),
				generateApptFromHoursAndMins(9, 0, 18, 0, 8, 8),
			},
			Error:           true,
			ExpectedError:   "strconv.ParseInt: parsing \"00000w\": invalid syntax",
			ExtendWeek:      false,
			ExtendDay:       false,
			InvalidBookings: false,
			InvalidProposed: true,
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.ScheduledAvailability)
			// Tests if 8 day week is ignored, by dropping the 8th day
			if tc.ExtendWeek {
				scheduled := append(*scheduledAvail, emptyDay)
				scheduledAvail = &scheduled
			}
			//  Tests if 25 hour day is ignored, by dropping the 25th hour
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)
			if tc.ExtendDay {
				(*scheduledAvail)[0] = (*scheduledAvail)[0] + "101010101010"
				(*bookings)[0] = (*bookings)[0] + "101010101010"
			}
			if tc.InvalidBookings {
				invalidBookings := (*bookings)[0][0:5] + "w"
				(*bookings)[0] = invalidBookings
			}

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)

			proposedAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.ProposedSchedule)
			if tc.InvalidProposed {
				invalidProposed := (*proposedAvail)[0][0:5] + "w"
				(*proposedAvail)[0] = invalidProposed
			}
			proposedSchedule := generateSchedule(*proposedAvail, *emptyWeek, &baseDate)

			computedSchedule, err := bScheduler.UpdateSchedule(proposedSchedule, schedule)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*proposedAvail)[i]
					computed := (*computedSchedule.Schedule)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestComposeAppointments(t *testing.T) {
	timeInterval := 5
	bScheduler, _ := NewBScheduler(timeInterval)

	apptStartTime, _ := time.Parse("2006-01-02T15:04:05Z", "2011-10-10T23:30:00Z")
	apptEndTime, _ := time.Parse("2006-01-02T15:04:05Z", "2011-10-11T00:30:00Z")

	apptToBook := &models.Appointment{
		StartTime: &apptStartTime,
		EndTime:   &apptEndTime,
	}

	expectedApptEndTime := time.Date(
		apptStartTime.Year(),
		apptStartTime.Month(),
		apptStartTime.Day(),
		23,
		59,
		59,
		0,
		time.UTC,
	)
	expectedAppt := &models.Appointment{
		StartTime: &apptStartTime,
		EndTime:   &expectedApptEndTime,
	}
	expectedSecondApptStartTime := time.Date(
		apptEndTime.Year(),
		apptEndTime.Month(),
		apptEndTime.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	expectedSecondAppt := &models.Appointment{
		StartTime: &expectedSecondApptStartTime,
		EndTime:   &apptEndTime,
	}

	apptDuo := bScheduler.ComposeAppointments(apptToBook)

	if !apptDuo.InitialAppointment.StartTime.Equal(*expectedAppt.StartTime) {
		t.Fatalf("expected initial StartTime: %s, received: %s", expectedAppt.StartTime.GoString(), apptDuo.InitialAppointment.StartTime.GoString())
	}
	if !apptDuo.InitialAppointment.EndTime.Equal(*expectedAppt.EndTime) {
		t.Fatalf("expected initial EndTime: %s, received: %s", expectedAppt.EndTime.GoString(), apptDuo.InitialAppointment.EndTime.GoString())
	}
	if !apptDuo.SecondAppointment.StartTime.Equal(*expectedSecondAppt.StartTime) {
		t.Fatalf("expected initial StartTime: %s, received: %s", expectedSecondAppt.StartTime.GoString(), apptDuo.SecondAppointment.StartTime.GoString())
	}
	if !apptDuo.SecondAppointment.EndTime.Equal(*expectedSecondAppt.EndTime) {
		t.Fatalf("expected initial EndTime: %s, received: %s", expectedSecondAppt.EndTime.GoString(), apptDuo.SecondAppointment.EndTime.GoString())
	}
}

func TestProcessAppointment(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-02T00:00:00Z")

	type test struct {
		Schedule      []models.Appointment
		Bookings      []models.Appointment
		Appt          models.Appointment
		ActionType    constants.ScheduleAction
		Error         bool
		Expected      []models.Appointment
		ExpectedError string
		Name          string
	}

	tests := []test{
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(10, 0, 10, 55, 3, 3),
			ActionType:    constants.BOOKING_UPDATE,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(10, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should call update appointments with only one appointment if the appointment does not cross the day boundary the type is delete",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(23, 30, 0, 30, 2, 3),
			ActionType:    constants.BOOKING_UPDATE,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(23, 30, 0, 30, 2, 3), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should call handleBookingUpdate with two appointments if the appointment crosses the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(10, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(10, 0, 11, 0, 3, 3),
			ActionType:    constants.DELETE_APPOINTMENT,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 5, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should call delete appointment with only one appointment if the appointment does not cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(23, 30, 0, 30, 2, 3),
			ActionType:    constants.DELETE_APPOINTMENT,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 25, 2, 2), generateApptFromHoursAndMins(0, 35, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should call delete appointment with two appointments if the appointment crosses the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(23, 30, 0, 30, 2, 3),
			ActionType:    constants.UNKNOWN,
			Error:         true,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 25, 2, 2), generateApptFromHoursAndMins(0, 35, 17, 0, 3, 3)},
			ExpectedError: "BScheduler Error: Recieved invalid action type: UNKNOWN",
			Name:          "should return false if passed an unknown action type",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			computedSchedule, err := bScheduler.ProcessAppointment(&tc.Appt, schedule, tc.ActionType)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestProcessAppointments(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-02T00:00:00Z")

	type test struct {
		Schedule      []models.Appointment
		Bookings      []models.Appointment
		Appts         []models.Appointment
		ActionType    constants.ScheduleAction
		Error         bool
		Expected      []models.Appointment
		ExpectedError string
		Name          string
	}

	tests := []test{
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(10, 0, 9, 0, 3, 3)},
			ActionType:    constants.BOOKING_UPDATE,
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BString Error: Appointment can't end before it begins.  Appointment start: time.Date(2020, time.February, 3, 10, 0, 0, 0, time.UTC) Appointment end: time.Date(2020, time.February, 3, 9, 0, 0, 0, time.UTC)",
			Name:          "should return throw an error if an appointment is invalid",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(10, 0, 12, 0, 3, 3), generateApptFromHoursAndMins(11, 0, 13, 0, 3, 3)},
			ActionType:    constants.BOOKING_UPDATE,
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BString Error: Appointment can't begin before previous appointment ends.  Appointment start: time.Date(2020, time.February, 3, 11, 0, 0, 0, time.UTC) Previous Appointment end: time.Date(2020, time.February, 3, 13, 0, 0, 0, time.UTC)",
			Name:          "should throw an error if the appointment array is invalid",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(10, 0, 11, 0, 3, 3)},
			ActionType:    constants.UNKNOWN,
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BScheduler Error: Recieved invalid action type: UNKNOWN",
			Name:          "should return an error if passed an unknown action type",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(11, 0, 11, 55, 3, 3), generateApptFromHoursAndMins(16, 0, 17, 0, 3, 3)},
			ActionType:    constants.DELETE_APPOINTMENT,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(12, 0, 15, 55, 3, 3)},
			ExpectedError: "",
			Name:          "should call deleteAppointments if the appropriate action type is passed",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(12, 0, 16, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(11, 0, 11, 55, 3, 3), generateApptFromHoursAndMins(16, 5, 17, 0, 3, 3)},
			ActionType:    constants.BOOKING_UPDATE,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should call handleBookingUpdateBString if the appropriate action type is passed",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			computedSchedule, err := bScheduler.ProcessAppointments(&tc.Appts, schedule, tc.ActionType)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestHandleBookingUpdate(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-02T00:00:00Z")

	type test struct {
		Schedule      []models.Appointment
		Bookings      []models.Appointment
		Appt          models.Appointment
		FirstAppt     *models.Appointment
		Error         bool
		Expected      []models.Appointment
		ExpectedError string
		Name          string
	}

	tests := []test{
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(10, 0, 10, 55, 3, 3),
			FirstAppt:     nil,
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(10, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that doesn't cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(12, 0, 2, 0, 3, 3),
			FirstAppt:     nil,
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BString Error: Appointment can't end before it begins.  Appointment start: time.Date(2020, time.February, 3, 12, 0, 0, 0, time.UTC) Appointment end: time.Date(2020, time.February, 3, 2, 0, 0, 0, time.UTC)",
			Name:          "should throw an error if the appointment passed is not valid",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(0, 0, 2, 0, 3, 3),
			FirstAppt:     nil,
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BScheduleUtil Error: Time intervals overlap.",
			Name:          "should return an error if the appointment passed does not fit in the schedule",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 22, 55, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(0, 0, 0, 55, 3, 3),
			FirstAppt:     generateApptFromHoursAndMinsPointer(23, 0, 23, 59, 2, 2),
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that does cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(0, 0, 1, 0, 3, 3),
			FirstAppt:     generateApptFromHoursAndMinsPointer(23, 59, 23, 0, 2, 2),
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BString Error: Appointment can't end before it begins.  Appointment start: time.Date(2020, time.February, 2, 23, 59, 0, 0, time.UTC) Appointment end: time.Date(2020, time.February, 2, 23, 0, 0, 0, time.UTC)",
			Name:          "should throw an error if a firstAppt is passed not a valid appointment",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appt:          generateApptFromHoursAndMins(0, 0, 1, 0, 3, 3),
			FirstAppt:     generateApptFromHoursAndMinsPointer(23, 0, 23, 59, 2, 2),
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BScheduleUtil Error: Time intervals overlap.",
			Name:          "should return an error if a firstAppt passed does not fit in the schedule",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			computedSchedule, err := bScheduler.HandleBookingUpdate(&tc.Appt, schedule, tc.FirstAppt)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestHandleBookingUpdateBString(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-02T00:00:00Z")

	type test struct {
		Schedule      []models.Appointment
		Bookings      []models.Appointment
		Appts         []models.Appointment
		Error         bool
		Expected      []models.Appointment
		ExpectedError string
		Name          string
	}

	tests := []test{
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(10, 0, 10, 55, 3, 3)},
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(10, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that doesn't cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(0, 0, 2, 0, 3, 3)},
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BScheduleUtil Error: Time intervals overlap.",
			Name:          "should return an error if the appointment passed does not fit in the schedule",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 22, 55, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(23, 0, 0, 55, 2, 3)},
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that does cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 0, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(23, 0, 1, 0, 2, 3)},
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BScheduleUtil Error: Time intervals overlap.",
			Name:          "should return an error if a firstAppt passed does not fit in the schedule",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 18, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 22, 55, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			Appts:         []models.Appointment{generateApptFromHoursAndMins(23, 0, 0, 55, 2, 3), generateApptFromHoursAndMins(17, 5, 18, 0, 3, 3)},
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 18, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle multiple appointments",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			apptsBString, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Appts)

			computedSchedule, err := bScheduler.HandleBookingUpdateBString(*apptsBString, schedule)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestBSchedulerDeleteAppointment(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	type test struct {
		Schedule          []models.Appointment
		Bookings          []models.Appointment
		ApptToDelete      models.Appointment
		FirstApptToDelete *models.Appointment
		BaseDate          string
		Error             bool
		Expected          []models.Appointment
		ExpectedError     string
		Name              string
	}

	tests := []test{
		{
			Schedule:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ApptToDelete:      generateApptFromHoursAndMins(11, 0, 11, 55, 3, 3),
			FirstApptToDelete: nil,
			BaseDate:          "2020-02-02T00:00:00Z",
			Error:             false,
			Expected:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(12, 0, 17, 0, 3, 3)},
			ExpectedError:     "",
			Name:              "should handle an appointment that doesn't cross the day boundary",
		},
		{
			Schedule:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ApptToDelete:      generateApptFromHoursAndMins(0, 0, 0, 55, 3, 3),
			FirstApptToDelete: generateApptFromHoursAndMinsPointer(23, 0, 23, 59, 2, 2),
			BaseDate:          "2020-02-02T00:00:00Z",
			Error:             false,
			Expected:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 22, 55, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			ExpectedError:     "",
			Name:              "should handle an appointment that does cross the day boundary",
		},
		{
			Schedule:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:          []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ApptToDelete:      generateApptFromHoursAndMins(9, 0, 11, 55, 3, 3),
			FirstApptToDelete: nil,
			BaseDate:          "2020-02-02T00:00:00Z",
			Error:             true,
			Expected:          []models.Appointment{},
			ExpectedError:     "error occured on hour: 9, error: BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 111111111111 Schedule: 000000000000",
			Name:              "should error if the deletion is invalid for one day",
		},
		{
			Schedule:          []models.Appointment{generateApptFromHoursAndMins(16, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:          []models.Appointment{generateApptFromHoursAndMins(16, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ApptToDelete:      generateApptFromHoursAndMins(0, 0, 0, 55, 3, 3),
			FirstApptToDelete: generateApptFromHoursAndMinsPointer(14, 0, 23, 59, 3, 3),
			BaseDate:          "2020-02-02T00:00:00Z",
			Error:             true,
			Expected:          []models.Appointment{},
			ExpectedError:     "error occured on hour: 17, error: BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 111111111111 Schedule: 100000000000",
			Name:              "should return false if the deletion is invalid for the first half",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)
			baseDate, _ := time.Parse("2006-01-02T15:04:05Z", tc.BaseDate)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			computedSchedule, err := bScheduler.DeleteAppointment(&tc.ApptToDelete, schedule, tc.FirstApptToDelete)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}

func TestDeleteAppointments(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)

	bScheduler, _ := NewBScheduler(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	type test struct {
		Schedule      []models.Appointment
		Bookings      []models.Appointment
		ApptsToDelete []models.Appointment
		BaseDate      string
		Error         bool
		Expected      []models.Appointment
		ExpectedError string
		Name          string
	}

	tests := []test{
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ApptsToDelete: []models.Appointment{generateApptFromHoursAndMins(11, 0, 11, 55, 3, 3)},
			BaseDate:      "2020-02-02T00:00:00Z",
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(12, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that doesn't cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ApptsToDelete: []models.Appointment{generateApptFromHoursAndMins(23, 0, 0, 55, 2, 3)},
			BaseDate:      "2020-02-02T00:00:00Z",
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 22, 55, 2, 2), generateApptFromHoursAndMins(1, 0, 17, 0, 3, 3)},
			ExpectedError: "",
			Name:          "should handle an appointment that does cross the day boundary",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ApptsToDelete: []models.Appointment{generateApptFromHoursAndMins(9, 0, 11, 55, 3, 3)},
			BaseDate:      "2020-02-02T00:00:00Z",
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BSchedule Error: interval to delete occurs outside of schedule on day 1 of the week starting on time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC)",
			Name:          "should return false if the deletion is invalid for one day",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(16, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(16, 0, 23, 59, 2, 2), generateApptFromHoursAndMins(0, 0, 17, 0, 3, 3)},
			ApptsToDelete: []models.Appointment{generateApptFromHoursAndMins(14, 0, 0, 55, 2, 3)},
			BaseDate:      "2020-02-02T00:00:00Z",
			Error:         true,
			Expected:      []models.Appointment{},
			ExpectedError: "BSchedule Error: interval to delete occurs outside of schedule on day 0 of the week starting on time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC)",
			Name:          "should throw an error if the deletion is invalid for the first half",
		},
		{
			Schedule:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(9, 0, 17, 0, 3, 3)},
			Bookings:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(11, 0, 17, 0, 3, 3)},
			ApptsToDelete: []models.Appointment{generateApptFromHoursAndMins(11, 0, 11, 55, 3, 3), generateApptFromHoursAndMins(16, 0, 17, 0, 3, 3)},
			BaseDate:      "2020-02-02T00:00:00Z",
			Error:         false,
			Expected:      []models.Appointment{generateApptFromHoursAndMins(8, 0, 18, 0, 2, 2), generateApptFromHoursAndMins(12, 0, 15, 55, 3, 3)},
			ExpectedError: "",
			Name:          "should handle multiple appointments",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			scheduledAvail, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Schedule)
			bookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Bookings)
			baseDate, _ := time.Parse("2006-01-02T15:04:05Z", tc.BaseDate)

			schedule := generateSchedule(*scheduledAvail, *bookings, &baseDate)
			apptsToDelete, _ := bStringUtil.GenerateBStringFromAppointments(&tc.ApptsToDelete)
			expectedBookings, _ := bStringUtil.GenerateBStringFromAppointments(&tc.Expected)

			computedSchedule, err := bScheduler.DeleteAppointments(*apptsToDelete, schedule)

			if !tc.Error && err != nil {
				t.Fatalf("expected no error, received: %s", err.Error())
			}
			if tc.Error && err.Error() != tc.ExpectedError {
				t.Fatalf("expected error with message: %s, received: %s", tc.ExpectedError, err.Error())
			}
			if !tc.Error {
				for i := 0; i < constants.DaysInWeek; i++ {
					expected := (*expectedBookings)[i]
					computed := (*computedSchedule.Bookings)[i]

					if strings.Compare(expected, computed) != 0 {
						t.Fatalf("bookings did not match expected on day: %d, expected: %s, received: %s", i, expected, computed)
					}
				}
			}
		})
	}
}
