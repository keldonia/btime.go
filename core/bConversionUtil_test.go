package core

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/keldonia/btime.go/constants"
	"github.com/keldonia/btime.go/models"
)

func generateMockDate(hour int, minute int, day *int) time.Time {
	if day == nil {
		one := 1
		day = &one
	}

	return time.Date(
		2020,
		2,
		*day,
		hour,
		minute,
		0,
		0,
		time.UTC,
	)
}

func generateEmptyDay(timeInterval int) string {
	intervalsInHour := constants.MinutesInHour / timeInterval
	emptyHour := strings.Repeat(constants.ZeroPad, intervalsInHour)
	emptyDay := strings.Repeat(emptyHour, constants.HoursInDay)

	return emptyDay
}

func generateEmptyWeek(timeInterval int) []string {
	emptyWeek := []string{}
	for i := 0; i < constants.DaysInWeek; i++ {
		emptyWeek = append(emptyWeek, generateEmptyDay(timeInterval))
	}

	return emptyWeek
}

func generateEmptyAppointmentWeek() *[][]models.Appointment {
	emptyAppointmentWeek := [][]models.Appointment{}

	for i := 0; i < constants.DaysInWeek; i++ {
		emptyAppointmentDay := generateEmptyAppointmentDay()
		emptyAppointmentWeek = append(emptyAppointmentWeek, *emptyAppointmentDay)
	}

	return &emptyAppointmentWeek
}

func generateEmptyAppointmentDay() *[]models.Appointment {
	return &[]models.Appointment{}
}

func generateAppoinmentsFromAppt(appt []models.Appointment) *[]models.Appointment {
	baseAppts := *generateEmptyAppointmentDay()
	baseAppts = append(baseAppts, appt...)

	return &baseAppts
}

func generateApptFromDateStrings(startStr string, endStr string) models.Appointment {
	start, _ := time.Parse("2006-01-02T15:04:05Z", startStr)
	end, _ := time.Parse("2006-01-02T15:04:05Z", endStr)

	return models.Appointment{
		StartTime: &start,
		EndTime:   &end,
	}
}

func TestConvertScheduleToAppointmentSchedule(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bConversionUtil, _ := NewBConversionUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-09T00:00:00Z")
	emptySchedule := generateEmptyWeek(timeInterval)
	emptyBookings := generateEmptyWeek(timeInterval)

	schedule := models.Schedule{
		Schedule:  &emptySchedule,
		Bookings:  &emptyBookings,
		WeekStart: &baseDate,
	}

	emptyAvail := generateEmptyWeek(timeInterval)

	expectedAppointmentSchedule := models.AppointmentSchedule{
		Schedule:     generateEmptyAppointmentWeek(),
		Bookings:     generateEmptyAppointmentWeek(),
		Availability: generateEmptyAppointmentWeek(),
		WeekStart:    &baseDate,
	}

	computedAppointmentSchedule := bConversionUtil.ConvertScheduleToAppointmentSchedule(&schedule, emptyAvail)

	var equal bool = true
	equal = len(*expectedAppointmentSchedule.Schedule) == len(*computedAppointmentSchedule.Schedule)
	if !equal {
		t.Fatalf("schedules are not equal, schedule1: %s, schedule2: %s", expectedAppointmentSchedule.Schedule, computedAppointmentSchedule.Schedule)
	}
	equal = len(*expectedAppointmentSchedule.Bookings) == len(*computedAppointmentSchedule.Bookings)
	if !equal {
		t.Fatalf("bookings are not equal")
	}
	equal = len(*expectedAppointmentSchedule.Availability) == len(*computedAppointmentSchedule.Availability)
	if !equal {
		t.Fatalf("availability are not equal")
	}
	equal = expectedAppointmentSchedule.WeekStart.Equal(*computedAppointmentSchedule.WeekStart)
	if !equal {
		t.Fatalf("weekStart are not equal")
	}
}

func TestCalculateDate(t *testing.T) {
	type test struct {
		Name        string
		InputDate   string
		TimePointer int
		Expected    string
		End         bool
	}

	tests := []test{
		{Name: "should properly calculate a datetime that is at the beginning of the day, and is the beginning of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 0, Expected: "2019-12-29T00:00:00Z", End: false},
		{Name: "should properly calculate a datetime that is at the beginning of the day, and is the end of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 0, Expected: "2019-12-29T00:04:59Z", End: true},
		{Name: "should properly calculate a datetime that is at the end of the day, and is the beginning of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 287, Expected: "2019-12-29T23:55:00Z", End: false},
		{Name: "should properly calculate a datetime that is at the end of the day, and is the end of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 287, Expected: "2019-12-29T23:59:59Z", End: true},
		{Name: "should properly calculate a datetime that is at the middle of the day, and is the beginning of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 144, Expected: "2019-12-29T12:00:00Z", End: false},
		{Name: "should properly calculate a datetime that is at the middle of the day, and is the end of an appointment", InputDate: "2019-12-29T00:00:00Z", TimePointer: 144, Expected: "2019-12-29T12:04:59Z", End: true},
	}

	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bConversionUtil, _ := NewBConversionUtil(bTimeConfig)

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			input, _ := time.Parse("2006-01-02T15:04:05Z", tc.InputDate)
			expected, _ := time.Parse("2006-01-02T15:04:05Z", tc.Expected)

			computed := bConversionUtil.CalculateDate(tc.TimePointer, &input, tc.End)

			if !expected.Equal(*computed) {
				t.Fatalf("expected: %s, recieved: %s", expected.GoString(), computed.GoString())
			}
		})
	}
}

func TestConvertTimeSlotsStringToAppointments(t *testing.T) {
	timeInterval := 5

	type test struct {
		Name      string
		TimeSlots models.Appointment
		Expected  *[]models.Appointment
	}
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bConversionUtil, _ := NewBConversionUtil(bTimeConfig)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2019-12-29T00:00:00Z")

	tests := []test{
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment that spans across the whole day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T00:00:00Z", "2019-12-29T23:59:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T00:00:00Z", "2019-12-29T23:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment in the middle of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:45:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:49:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment at the start of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:49:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:49:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment close to the start of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T00:05:00Z", "2019-12-29T14:49:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T00:05:00Z", "2019-12-29T14:49:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment at the end of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment close to the end of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:54:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:54:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment close to the end of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment close to the end of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with one appointment, if there was one contigous segment close to the end of the day",
			TimeSlots: generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z"),
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T23:59:59Z")}),
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			timeSlotsStr, err := bStringUtil.GenerateBString(&tc.TimeSlots)
			if err != nil {
				fmt.Println(err)
			}
			computed := bConversionUtil.ConvertTimeSlotsStringToAppointments(*timeSlotsStr, &baseDate)

			expectedAppt := (*tc.Expected)[0]
			computedAppt := (*computed)[0]

			failed := len(*computed) != len(*tc.Expected)
			if failed {
				t.Fatalf("did not receive expected appts, different lengths expected: %d received: %d", len(*tc.Expected), len(*computed))
			}

			failed = !computedAppt.StartTime.Equal(*expectedAppt.StartTime)
			if failed {
				t.Fatalf("did not receive expected appt, expected start: %s received start %s", expectedAppt.StartTime.GoString(), computedAppt.StartTime.GoString())
			}

			failed = !computedAppt.EndTime.Equal(*expectedAppt.EndTime)
			if failed {
				t.Fatalf("did not receive expected appt, expected end: %s received end %s", expectedAppt.EndTime.GoString(), computedAppt.EndTime.GoString())
			}
		})
	}

	type test2 struct {
		Name      string
		TimeSlots string
		Expected  *[]models.Appointment
	}

	tests2 := []test2{
		{
			Name:      "should return an empty appointment array, if there were no time slots",
			TimeSlots: generateEmptyDay(timeInterval),
			Expected:  generateEmptyAppointmentDay(),
		},
		{
			Name:      "should return an appointment array with two appointment, if there were two segments",
			TimeSlots: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000111111111111111111111111000000000000111111111111111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T10:00:00Z", "2019-12-29T11:59:59Z"), generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with three appointment, if there were three segments",
			TimeSlots: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000111111111111111111111111000000000000111111111111111111111111000000000000111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T10:00:00Z", "2019-12-29T11:59:59Z"), generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:59:59Z"), generateApptFromDateStrings("2019-12-29T16:00:00Z", "2019-12-29T16:59:59Z")}),
		},
		{
			Name:      "should return an appointment array with three appointment, if there were three segments, should ignore excess intervals",
			TimeSlots: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000111111111111111111111111000000000000111111111111111111111111000000000000111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000111111",
			Expected:  generateAppoinmentsFromAppt([]models.Appointment{generateApptFromDateStrings("2019-12-29T10:00:00Z", "2019-12-29T11:59:59Z"), generateApptFromDateStrings("2019-12-29T13:00:00Z", "2019-12-29T14:59:59Z"), generateApptFromDateStrings("2019-12-29T16:00:00Z", "2019-12-29T16:59:59Z")}),
		},
	}

	for i := 0; i < len(tests2); i++ {
		tc := tests2[i]
		t.Run(tc.Name, func(t *testing.T) {
			computed := bConversionUtil.ConvertTimeSlotsStringToAppointments(tc.TimeSlots, &baseDate)

			if len(*tc.Expected) != len(*computed) {
				t.Fatal("recieved a different number of appts than expected")
			}

			for i := 0; i < len(*tc.Expected); i++ {
				expectedAppt := (*tc.Expected)[i]
				computedAppt := (*computed)[i]

				if !expectedAppt.StartTime.Equal(*computedAppt.StartTime) || !expectedAppt.EndTime.Equal(*computedAppt.EndTime) {
					t.Fatalf("computed appointment %d did not match the expected appt", i)
				}
			}
		})
	}
}
