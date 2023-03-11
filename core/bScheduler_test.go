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

func TestNewBSchedulerinvalidInput(t *testing.T) {
	timeInterval := -1
	expectedErrorStr := fmt.Sprintf("[BConfig] received an invalid time interval: %d", timeInterval)

	bScheduler, err := NewBScheduler(timeInterval)

	if err.Error() != expectedErrorStr {
		t.Fatalf("expected error: %s, received: %s", expectedErrorStr, err.Error())
	}

	if bScheduler != nil {
		t.Fatalf("expected bScheduler to not be generated")
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
