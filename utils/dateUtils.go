package utils

import (
	"time"

	"github.com/keldonia/btime.go/models"
)

// Utility function to ensure both Dates in and appointment are UTC, converting to UTC if not
func EnforceUTC(appt *models.Appointment) *models.Appointment {
	return &models.Appointment{
		StartTime: GetUTC(*appt.StartTime),
		EndTime:   GetUTC(*appt.EndTime),
	}
}

// Utility function to ensure a Date is UTC
func GetUTC(date time.Time) *time.Time {
	utcTime := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		0,
		0,
		time.UTC,
	)

	return &utcTime
}

// Takes in a date and gets the UTC date and returns the UTC date that began the week that date is in
func GetFirstDayOfWeekFromDate(date *time.Time) *time.Time {
	startOfWeek := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
		0,
		0,
		time.UTC,
	)

	weekdayOffset := -1 * int(date.Weekday())

	firstDayOfWeek := startOfWeek.AddDate(0, 0, weekdayOffset)

	return &firstDayOfWeek
}

// Takes a date and generates a base Date for each day of the week
func GetDatesFromStartDate(date *time.Time) []time.Time {
	returnDates := []time.Time{*date}

	// NB: We only need to create a Date for each day of the week.
	// Additionally, Date::UTC automatically rolls over to the next
	// largest increment, if a value is greater than the max
	// day 0 = Sunday
	for i := 1; i < 7; i++ {
		modifiedDate := date.AddDate(0, 0, i)

		returnDates = append(returnDates, modifiedDate)
	}

	return returnDates
}

// Takes an appointment and checks if the appoint crosses a day boundry
//
// NB: We assume that at most appts cross 1 day boundary
func CrossesDayBoundary(appt models.Appointment) bool {
	return appt.StartTime.UTC().YearDay() != appt.EndTime.UTC().YearDay()
}

// Takes an appointment and checks if the appoint crosses a week boundry
func CrossesWeekBoundary(appt models.Appointment) bool {
	return GetWeek(*appt.StartTime) != GetWeek(*appt.EndTime)
}

// Takes date and returns the ISO Week
func GetWeek(date time.Time) int {
	_, weekNumber := date.UTC().ISOWeek()

	return weekNumber
}

// Takes a date and returns the utc start of the day
func GetUTCDateStart(date time.Time) *time.Time {
	utcStartDate := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	return &utcStartDate
}

// Takes a date and returns the utc end of the day
func GetUTCDateEnd(date *time.Time, seconds int) *time.Time {
	utcEndDate := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		23,
		59,
		59,
		0,
		time.UTC,
	)

	return &utcEndDate
}
