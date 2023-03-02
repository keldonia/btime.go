package utils

import (
	"time"

	"github.com/keldonia/btime.go/models"
)

func EnforceUTC(appt *models.Appointment) *models.Appointment {
	return &models.Appointment{
		StartTime: GetUTC(*appt.StartTime),
		EndTime:   GetUTC(*appt.EndTime),
	}
}

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

func GetDatesFromStartDate(date *time.Time) []time.Time {
	returnDates := []time.Time{*date}

	for i := 1; i < 7; i++ {
		modifiedDate := date.AddDate(0, 0, i)

		returnDates = append(returnDates, modifiedDate)
	}

	return returnDates
}

func CrossesDayBoundary(appt models.Appointment) bool {
	return appt.StartTime.UTC().YearDay() != appt.EndTime.UTC().YearDay()
}

func CrossesWeekBoundary(appt models.Appointment) bool {
	return GetWeek(*appt.StartTime) != GetWeek(*appt.EndTime)
}

func GetWeek(date time.Time) int {
	_, weekNumber := date.UTC().ISOWeek()

	return weekNumber
}

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
