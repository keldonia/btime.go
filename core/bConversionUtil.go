package core

import (
	"fmt"
	"math"
	"time"

	"github.com/keldonia/btime.go/models"
	"github.com/keldonia/btime.go/utils"
)

//go:generate mockery --name BConversionUtil
type BConversionUtil interface {
	ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule
	ConvertTimeSlotsStringToAppointments(timeSlots string, date *time.Time) *[]models.Appointment
	CalculateDate(timePointerIndex int, baseDate *time.Time, end bool) *time.Time
}

type BConversionUtilImpl struct {
	bTimeConfig *BTimeConfig
}

// Instantiates a new BinaryConversionUtil, which is responsible for handling the conversion of schedules to Appointments
func NewBConversionUtil(bTimeConfig *BTimeConfig) (BConversionUtil, error) {
	if bTimeConfig == nil {
		return nil, fmt.Errorf("[BConversionUtil] No BTimeConfig was provided")
	}

	return &BConversionUtilImpl{
		bTimeConfig: bTimeConfig,
	}, nil
}

// Takes a schedule and the schedule's remaining availability,
// and converts each of the bTime representations into Appointment arrays
func (bcu *BConversionUtilImpl) ConvertScheduleToAppointmentSchedule(schedule *models.Schedule, availability []string) *models.AppointmentSchedule {
	days := utils.GetDatesFromStartDate(schedule.WeekStart)
	appointmentAvailability := [][]models.Appointment{}
	appointmentBookings := [][]models.Appointment{}
	appointmentBaseSchedule := [][]models.Appointment{}

	for i := 0; i < len(availability); i++ {
		avail := availability[i]
		day := days[i]
		appointments := bcu.ConvertTimeSlotsStringToAppointments(avail, &day)
		appointmentAvailability = append(appointmentAvailability, *appointments)
	}

	for i := 0; i < len(*schedule.Bookings); i++ {
		booking := (*schedule.Bookings)[i]
		day := days[i]
		bookingSet := bcu.ConvertTimeSlotsStringToAppointments(booking, &day)
		appointmentBookings = append(appointmentBookings, *bookingSet)
	}

	for i := 0; i < len(*schedule.Schedule); i++ {
		daySchedule := (*schedule.Schedule)[i]
		day := days[i]
		dayApptSchedule := bcu.ConvertTimeSlotsStringToAppointments(daySchedule, &day)
		appointmentBaseSchedule = append(appointmentBaseSchedule, *dayApptSchedule)
	}

	return &models.AppointmentSchedule{
		WeekStart:    schedule.WeekStart,
		Bookings:     &appointmentBookings,
		Availability: &appointmentAvailability,
		Schedule:     &appointmentBaseSchedule,
	}
}

// Takes a set of timeslots and the date on which they occurred and converts them into Appointments
func (bcu *BConversionUtilImpl) ConvertTimeSlotsStringToAppointments(timeSlots string, date *time.Time) *[]models.Appointment {
	appointments := []models.Appointment{}
	var currentStart *time.Time

	for i := 0; i < bcu.bTimeConfig.IntervalsInDay; i++ {
		if string(timeSlots[i]) == "1" && currentStart == nil {
			currentStart = bcu.CalculateDate(i, date, false)
		}
		if currentStart != nil && string(timeSlots[i]) == "0" {
			currentEnd := bcu.CalculateDate(i-1, date, true)
			appointment := models.Appointment{
				StartTime: currentStart,
				EndTime:   currentEnd,
			}

			appointments = append(appointments, appointment)
			currentStart = nil
		}
	}

	if currentStart != nil {
		currentEnd := utils.GetUTCDateEnd(date, 59)
		appointment := models.Appointment{
			StartTime: currentStart,
			EndTime:   currentEnd,
		}
		appointments = append(appointments, appointment)
	}

	return &appointments
}

// Takes a  time pointer, base date — the date on which it occured,
// and boolean if it is the end of an appointmen and converts it into a Date
func (bcu *BConversionUtilImpl) CalculateDate(timePointerIndex int, baseDate *time.Time, end bool) *time.Time {
	hourMarker := float64(timePointerIndex / bcu.bTimeConfig.IntervalsInHour)
	hours := int(math.Floor(hourMarker))
	minutes := timePointerIndex % bcu.bTimeConfig.IntervalsInHour * bcu.bTimeConfig.TimeInterval
	seconds := 0

	if end {
		minutes += bcu.bTimeConfig.TimeInterval
		seconds = -1
	}

	returnTime := time.Date(
		baseDate.Year(),
		baseDate.Month(),
		baseDate.Day(),
		hours,
		minutes,
		seconds,
		0,
		time.UTC,
	)

	return &returnTime
}
