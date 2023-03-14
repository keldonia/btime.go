package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/keldonia/btime.go/models"
)

func TestEnforceUTC(t *testing.T) {
	startTime, _ := time.Parse("2006-01-02T15:04:05Z", "2011-10-10T23:30:00Z")
	endTime, _ := time.Parse("2006-01-02T15:04:05Z", "2011-10-10T23:30:00Z")

	apptToBook := models.Appointment{
		StartTime: &startTime,
		EndTime:   &endTime,
	}
	expectedStartTime := time.Date(
		startTime.Year(),
		startTime.Month(),
		startTime.Day(),
		startTime.Hour(),
		startTime.Minute(),
		0,
		0,
		time.UTC,
	)
	expectedEndTime := time.Date(
		endTime.Year(),
		endTime.Month(),
		endTime.Day(),
		endTime.Hour(),
		endTime.Minute(),
		0,
		0,
		time.UTC,
	)

	expectedAppt := models.Appointment{
		StartTime: &expectedStartTime,
		EndTime:   &expectedEndTime,
	}

	computedUtcAppt := EnforceUTC(&apptToBook)

	if !expectedAppt.StartTime.Equal(*computedUtcAppt.StartTime) || !expectedAppt.EndTime.Equal(*computedUtcAppt.EndTime) {
		t.Fatalf("expected start: %s, end: %s, received start: %s, end: %s", expectedAppt.StartTime.GoString(), expectedAppt.EndTime.GoString(), computedUtcAppt.StartTime.GoString(), computedUtcAppt.EndTime.GoString())
	}
}

func TestGetFirstDayOfWeekFromDate(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "2020-02-12T00:00:00Z", expected: "2020-02-09T00:00:00Z"},
		{input: "2020-02-09T00:00:00Z", expected: "2020-02-09T00:00:00Z"},
		{input: "2020-02-15T00:00:00Z", expected: "2020-02-09T00:00:00Z"},
		{input: "2020-02-11T00:00:00Z", expected: "2020-02-09T00:00:00Z"},
		{input: "2020-04-01T00:00:00Z", expected: "2020-03-29T00:00:00Z"},
		{input: "2020-04-04T00:00:00Z", expected: "2020-03-29T00:00:00Z"},
		{input: "2020-03-29T00:00:00Z", expected: "2020-03-29T00:00:00Z"},
		{input: "2019-12-30T00:00:00Z", expected: "2019-12-29T00:00:00Z"},
		{input: "2020-01-01T00:00:00Z", expected: "2019-12-29T00:00:00Z"},
		{input: "2020-01-03T00:00:00Z", expected: "2019-12-29T00:00:00Z"},
	}

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d:", i), func(t *testing.T) {
			input, _ := time.Parse("2006-01-02T15:04:05Z", tc.input)
			expected, _ := time.Parse("2006-01-02T15:04:05Z", tc.expected)

			computed := GetFirstDayOfWeekFromDate(&input)

			if !expected.Equal(*computed) {
				t.Fatalf("expected: %s, for input: %s, received: %s", expected.UTC().GoString(), input.UTC().GoString(), computed.UTC().GoString())
			}
		})
	}
}

func TestGetDatesFromStartDate(t *testing.T) {
	type test struct {
		Name      string
		StartDate string
		DayOne    string
		DayTwo    string
		DayThree  string
		DayFour   string
		DayFive   string
		DaySix    string
	}

	tests := []test{
		{
			Name:      "should properly create a weeks worth of Dates from a schedule, contained within a month",
			StartDate: "2020-02-09T00:00:00Z",
			DayOne:    "2020-02-10T00:00:00Z",
			DayTwo:    "2020-02-11T00:00:00Z",
			DayThree:  "2020-02-12T00:00:00Z",
			DayFour:   "2020-02-13T00:00:00Z",
			DayFive:   "2020-02-14T00:00:00Z",
			DaySix:    "2020-02-15T00:00:00Z",
		},
		{
			Name:      "should properly create a weeks worth of Dates from a schedule, that crosses a month boundary",
			StartDate: "2020-03-29T00:00:00Z",
			DayOne:    "2020-03-30T00:00:00Z",
			DayTwo:    "2020-03-31T00:00:00Z",
			DayThree:  "2020-04-01T00:00:00Z",
			DayFour:   "2020-04-02T00:00:00Z",
			DayFive:   "2020-04-03T00:00:00Z",
			DaySix:    "2020-04-04T00:00:00Z",
		},
		{
			Name:      "should properly create a weeks worth of Dates from a schedule, that crosses a year boundary",
			StartDate: "2019-12-29T00:00:00Z",
			DayOne:    "2019-12-30T00:00:00Z",
			DayTwo:    "2019-12-31T00:00:00Z",
			DayThree:  "2020-01-01T00:00:00Z",
			DayFour:   "2020-01-02T00:00:00Z",
			DayFive:   "2020-01-03T00:00:00Z",
			DaySix:    "2020-01-04T00:00:00Z",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			startDate, _ := time.Parse("2006-01-02T15:04:05Z", tc.StartDate)
			dayOne, _ := time.Parse("2006-01-02T15:04:05Z", tc.DayOne)
			dayTwo, _ := time.Parse("2006-01-02T15:04:05Z", tc.DayTwo)
			dayThree, _ := time.Parse("2006-01-02T15:04:05Z", tc.DayThree)
			dayFour, _ := time.Parse("2006-01-02T15:04:05Z", tc.DayFour)
			dayFive, _ := time.Parse("2006-01-02T15:04:05Z", tc.DayFive)
			daySix, _ := time.Parse("2006-01-02T15:04:05Z", tc.DaySix)

			expectedWeek := []time.Time{
				startDate,
				dayOne,
				dayTwo,
				dayThree,
				dayFour,
				dayFive,
				daySix,
			}

			computedWeek := GetDatesFromStartDate(&startDate)

			if len(computedWeek) != 7 {
				t.Fatalf("expected a week of 7 days, recieved %d days", len(computedWeek))
			}

			for i := 0; i < len(expectedWeek); i++ {
				expected := expectedWeek[i]
				computed := computedWeek[i]
				if !expected.Equal(computed) {
					t.Fatalf("day: %d unequal, expected: %s, recieved: %s", i, expected, computed)
				}
			}
		})
	}
}

func TestCrossesDayBoundary(t *testing.T) {
	type test struct {
		Name      string
		StartTime string
		Expected  bool
	}

	tests := []test{
		{
			Name:      "returns false if an appointment does not cross the day boundary",
			StartTime: "2011-10-10T10:48:00Z",
			Expected:  false,
		},
		{
			Name:      "returns true if an appointment does cross the day boundary",
			StartTime: "2011-10-10T23:48:00Z",
			Expected:  true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			startTime, _ := time.Parse("2006-01-02T15:04:05Z", tc.StartTime)
			endTime := startTime.Add(time.Hour)

			appt := models.Appointment{
				StartTime: &startTime,
				EndTime:   &endTime,
			}

			computed := CrossesDayBoundary(appt)

			if tc.Expected != computed {
				t.Fatalf("expected: %t, recieved: %t", tc.Expected, computed)
			}
		})
	}
}

func TestCrossesWeekBoundary(t *testing.T) {
	type test struct {
		DateOne  string
		DateTwo  string
		Expected bool
	}

	tests := []test{
		{DateOne: "1969-12-21T00:00:00Z", DateTwo: "1969-12-21T00:00:00Z", Expected: false},
		{DateOne: "1969-12-21T00:00:00Z", DateTwo: "1969-12-26T00:00:00Z", Expected: false},
		{DateOne: "1969-12-21T00:00:00Z", DateTwo: "1969-12-28T00:00:00Z", Expected: true},
		{DateOne: "1969-12-28T00:00:00Z", DateTwo: "1970-01-02T00:00:00Z", Expected: false},
		{DateOne: "2020-02-05T00:00:00Z", DateTwo: "2020-02-08T00:00:00Z", Expected: false},
		{DateOne: "2020-02-08T00:00:00Z", DateTwo: "2020-03-08T00:00:00Z", Expected: true},
		{DateOne: "2020-03-10T00:00:00Z", DateTwo: "2020-03-13T00:00:00Z", Expected: false},
		{DateOne: "2020-02-04T00:00:00Z", DateTwo: "2020-02-08T00:00:00Z", Expected: false},
		{DateOne: "2020-02-05T00:00:00Z", DateTwo: "2020-02-08T00:00:00Z", Expected: false},
	}

	for _, tc := range tests {
		tc := tc
		name := fmt.Sprintf("should return %t when passed an appt string starting at: %s and ending at %s", tc.Expected, tc.DateOne, tc.DateTwo)
		t.Run(name, func(t *testing.T) {
			startTime, _ := time.Parse("2006-01-02T15:04:05Z", tc.DateOne)
			endTime, _ := time.Parse("2006-01-02T15:04:05Z", tc.DateTwo)

			appt := models.Appointment{
				StartTime: &startTime,
				EndTime:   &endTime,
			}

			computed := CrossesWeekBoundary(appt)

			if tc.Expected != computed {
				t.Fatalf("expected: %t, recieved: %t", tc.Expected, computed)
			}
		})
	}
}

func TestGetWeek(t *testing.T) {
	type test struct {
		Input    string
		Expected int
	}

	tests := []test{
		{Input: "1969-12-21T00:00:00Z", Expected: 52},
		{Input: "1969-12-22T00:00:00Z", Expected: 52},
		{Input: "1969-12-23T00:00:00Z", Expected: 52},
		{Input: "1969-12-24T00:00:00Z", Expected: 52},
		{Input: "1969-12-25T00:00:00Z", Expected: 52},
		{Input: "1969-12-26T00:00:00Z", Expected: 52},
		{Input: "1969-12-27T00:00:00Z", Expected: 52},
		{Input: "1969-12-28T00:00:00Z", Expected: 1},
		{Input: "1969-12-29T00:00:00Z", Expected: 1},
		{Input: "1969-12-30T00:00:00Z", Expected: 1},
		{Input: "1969-12-31T00:00:00Z", Expected: 1},
		{Input: "1970-01-01T00:00:00Z", Expected: 1},
		{Input: "1970-01-02T00:00:00Z", Expected: 1},
		{Input: "1970-01-03T00:00:00Z", Expected: 1},
		{Input: "1970-01-04T00:00:00Z", Expected: 2},
		{Input: "1970-01-05T00:00:00Z", Expected: 2},
		{Input: "1970-01-06T00:00:00Z", Expected: 2},
		{Input: "1970-01-07T00:00:00Z", Expected: 2},
		{Input: "1970-01-08T00:00:00Z", Expected: 2},
		{Input: "1970-01-09T00:00:00Z", Expected: 2},
		{Input: "1970-01-10T00:00:00Z", Expected: 2},
		{Input: "2020-02-02T00:00:00Z", Expected: 6},
		{Input: "2020-02-03T00:00:00Z", Expected: 6},
		{Input: "2020-02-04T00:00:00Z", Expected: 6},
		{Input: "2020-02-05T00:00:00Z", Expected: 6},
		{Input: "2020-02-06T00:00:00Z", Expected: 6},
		{Input: "2020-02-07T00:00:00Z", Expected: 6},
		{Input: "2020-02-08T00:00:00Z", Expected: 6},
		{Input: "2020-03-08T00:00:00Z", Expected: 11},
		{Input: "2020-03-09T00:00:00Z", Expected: 11},
		{Input: "2020-03-10T00:00:00Z", Expected: 11},
		{Input: "2020-03-11T00:00:00Z", Expected: 11},
		{Input: "2020-03-12T00:00:00Z", Expected: 11},
		{Input: "2020-03-13T00:00:00Z", Expected: 11},
		{Input: "2020-03-14T00:00:00Z", Expected: 11},
	}

	for _, tc := range tests {
		tc := tc
		name := fmt.Sprintf("should return %d, when passed %s", tc.Expected, tc.Input)
		t.Run(name, func(t *testing.T) {
			date, _ := time.Parse("2006-01-02T15:04:05Z", tc.Input)

			computed := GetWeek(date)

			if tc.Expected != computed {
				t.Fatalf("expected: %d, recieved: %d", tc.Expected, computed)
			}
		})
	}
}

func TestGetUtcDateStart(t *testing.T) {
	initialDate, _ := time.Parse("2006-01-02T15:04:05Z", "1969-12-21T12:00:00Z")
	expectedDate, _ := time.Parse("2006-01-02T15:04:05Z", "1969-12-21T00:00:00Z")

	calculatedDate := GetUTCDateStart(initialDate)

	if !expectedDate.Equal(*calculatedDate) {
		t.Fatalf("expected: %s recieved: %s", expectedDate.GoString(), initialDate.GoString())
	}
}

func TestGetUtcDateEnd(t *testing.T) {
	type test struct {
		Name      string
		InputTime string
		Seconds   int
		Expected  string
	}

	tests := []test{
		{Name: "should return the expected utc end of the day", InputTime: "1969-12-21T12:00:00Z", Seconds: 0, Expected: "1969-12-21T23:59:00Z"},
		{Name: "should return the expected utc end of the day, with seconds", InputTime: "1969-12-21T12:00:00Z", Seconds: 59, Expected: "1969-12-21T23:59:59Z"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			input, _ := time.Parse("2006-01-02T15:04:05Z", tc.InputTime)
			expected, _ := time.Parse("2006-01-02T15:04:05Z", tc.Expected)

			computed := GetUTCDateEnd(&input, tc.Seconds)

			if !expected.Equal(*computed) {
				t.Fatalf("expected: %s, recieved: %s", expected.GoString(), computed.GoString())
			}
		})
	}
}
