package core

import (
	"fmt"
	"testing"
	"time"
)

func TestBadBPointerCalculatorSetup(t *testing.T) {
	bPointerCalculator, err := NewBPointerCalculator(nil)

	if bPointerCalculator != nil {
		t.Fatalf("expected bPointerCalculator to be nil")
	}

	if err.Error() != "[BPointerCalculator] No BTimeConfig was provided" {
		t.Fatalf("received an unexpected error: %s", err.Error())
	}
}

func TestFindBPointerIncludingDay5MinInterval(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bPointerCalculator, _ := NewBPointerCalculator(bTimeConfig)

	type test struct {
		Hour     int
		Minute   int
		Day      int
		Expected int
	}

	// NB: Feb 2, 2020 is a Sunday
	tests := []test{
		{Hour: 0, Minute: 0, Day: 2, Expected: 0 + 0*bTimeConfig.IntervalsInDay},
		{Hour: 0, Minute: 4, Day: 3, Expected: 1 * bTimeConfig.IntervalsInDay},
		{Hour: 0, Minute: 5, Day: 3, Expected: 1 + bTimeConfig.IntervalsInDay},
		{Hour: 1, Minute: 0, Day: 4, Expected: 12 + 2*bTimeConfig.IntervalsInDay},
		{Hour: 1, Minute: 1, Day: 5, Expected: 12 + 3*bTimeConfig.IntervalsInDay},
		{Hour: 0, Minute: 47, Day: 4, Expected: 9 + 2*bTimeConfig.IntervalsInDay},
		{Hour: 0, Minute: 5, Day: 8, Expected: 1 + 6*bTimeConfig.IntervalsInDay},
		{Hour: 12, Minute: 0, Day: 7, Expected: 144 + 5*bTimeConfig.IntervalsInDay},
		{Hour: 13, Minute: 31, Day: 6, Expected: 162 + 4*bTimeConfig.IntervalsInDay},
		{Hour: 5, Minute: 25, Day: 7, Expected: 65 + 5*bTimeConfig.IntervalsInDay},
		{Hour: 8, Minute: 15, Day: 5, Expected: 99 + 3*bTimeConfig.IntervalsInDay},
		{Hour: 10, Minute: 42, Day: 2, Expected: 128 + 0*bTimeConfig.IntervalsInDay},
		{Hour: 20, Minute: 7, Day: 6, Expected: 241 + 4*bTimeConfig.IntervalsInDay},
		{Hour: 23, Minute: 59, Day: 8, Expected: 287 + 6*bTimeConfig.IntervalsInDay},
		{Hour: 24, Minute: 0, Day: 2, Expected: 288},
		{Hour: 9, Minute: 0, Day: 6, Expected: 108 + 4*bTimeConfig.IntervalsInDay},
		{Hour: 12, Minute: 0, Day: 4, Expected: 144 + 2*bTimeConfig.IntervalsInDay},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("should return %d if hour is %d, minute is %d, and the day of the week is %d", tc.Expected, tc.Hour, tc.Minute, tc.Day)
		t.Run(name, func(t *testing.T) {
			date := generateMockDate(tc.Hour, tc.Minute, &tc.Day)
			bPointer := bPointerCalculator.FindBPointerIncludingDay(&date)

			if bPointer != tc.Expected {
				t.Fatalf("for date: %s expected: %d, recieved: %d", date.GoString(), tc.Expected, bPointer)
			}
		})
	}
}

func TestFindBPointerModiferForDayOfWeek(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bPointerCalculator, _ := NewBPointerCalculator(bTimeConfig)

	type test struct {
		Date     string
		Expected int
	}

	tests := []test{
		{Date: "2020-02-09T00:00:00Z", Expected: 0 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-10T00:00:00Z", Expected: 1 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-11T00:00:00Z", Expected: 2 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-12T00:00:00Z", Expected: 3 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-13T00:00:00Z", Expected: 4 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-14T00:00:00Z", Expected: 5 * bTimeConfig.IntervalsInDay},
		{Date: "2020-02-15T00:00:00Z", Expected: 6 * bTimeConfig.IntervalsInDay},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		dateTime, _ := time.Parse("2006-01-02T15:04:05Z", tc.Date)
		name := fmt.Sprintf("should return %d when passed %s, which has is %s", tc.Expected, tc.Date, dateTime.Weekday())
		t.Run(name, func(t *testing.T) {
			bPointer := bPointerCalculator.FindBPointerModiferForDayOfWeek(&dateTime)

			if bPointer != tc.Expected {
				t.Fatalf("for date: %s expected: %d, recieved: %d", dateTime.GoString(), tc.Expected, bPointer)
			}
		})
	}
}

func TestFindBPointer(t *testing.T) {
	type test struct {
		Hour     int
		Minute   int
		Expected int
	}

	one := 1

	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bPointerCalculator, _ := NewBPointerCalculator(bTimeConfig)

	tests := []test{
		{Hour: 0, Minute: 0, Expected: 0},
		{Hour: 0, Minute: 4, Expected: 0},
		{Hour: 0, Minute: 5, Expected: 1},
		{Hour: 1, Minute: 0, Expected: 12},
		{Hour: 1, Minute: 1, Expected: 12},
		{Hour: 0, Minute: 47, Expected: 9},
		{Hour: 0, Minute: 5, Expected: 1},
		{Hour: 12, Minute: 0, Expected: 144},
		{Hour: 13, Minute: 31, Expected: 162},
		{Hour: 5, Minute: 25, Expected: 65},
		{Hour: 8, Minute: 15, Expected: 99},
		{Hour: 10, Minute: 42, Expected: 128},
		{Hour: 20, Minute: 7, Expected: 241},
		{Hour: 23, Minute: 59, Expected: 287},
		{Hour: 24, Minute: 0, Expected: 0},
		{Hour: 9, Minute: 0, Expected: 108},
		{Hour: 12, Minute: 0, Expected: 144},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("should return %d if hour is %d and minute is %d", tc.Expected, tc.Hour, tc.Minute)
		t.Run(name, func(t *testing.T) {
			date := generateMockDate(tc.Hour, tc.Minute, &one)
			bPointer := bPointerCalculator.FindBPointer(&date)

			if bPointer != tc.Expected {
				t.Fatalf("for date: %s expected: %d, recieved: %d", date.GoString(), tc.Expected, bPointer)
			}
		})
	}
}
