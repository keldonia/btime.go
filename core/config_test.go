package core

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/keldonia/btime.go/constants"
)

func TestInvalidTimeSet(t *testing.T) {
	bTimeConfig, err := BuildConfigFromTimeInterval(-1)

	if bTimeConfig != nil {
		t.Fatalf("expected bTimeConfig to be nil")
	}

	if err.Error() != "[BConfig] received an invalid time interval: -1" {
		t.Fatalf("received an unexpected error: %s", err.Error())
	}
}

func TestConfigProperlyGenerates(t *testing.T) {
	type test struct {
		TimeInterval    int
		IntervalsInHour int
		IntervalsInDay  int
		IntervalsInWeek int
	}

	tests := []test{
		{TimeInterval: 1, IntervalsInHour: 60, IntervalsInDay: 1440, IntervalsInWeek: 10080},
		{TimeInterval: 2, IntervalsInHour: 30, IntervalsInDay: 720, IntervalsInWeek: 5040},
		{TimeInterval: 3, IntervalsInHour: 20, IntervalsInDay: 480, IntervalsInWeek: 3360},
		{TimeInterval: 4, IntervalsInHour: 15, IntervalsInDay: 360, IntervalsInWeek: 2520},
		{TimeInterval: 5, IntervalsInHour: 12, IntervalsInDay: 288, IntervalsInWeek: 2016},
		{TimeInterval: 6, IntervalsInHour: 10, IntervalsInDay: 240, IntervalsInWeek: 1680},
		{TimeInterval: 10, IntervalsInHour: 6, IntervalsInDay: 144, IntervalsInWeek: 1008},
		{TimeInterval: 12, IntervalsInHour: 5, IntervalsInDay: 120, IntervalsInWeek: 840},
		{TimeInterval: 15, IntervalsInHour: 4, IntervalsInDay: 96, IntervalsInWeek: 672},
		{TimeInterval: 20, IntervalsInHour: 3, IntervalsInDay: 72, IntervalsInWeek: 504},
		{TimeInterval: 30, IntervalsInHour: 2, IntervalsInDay: 48, IntervalsInWeek: 336},
		{TimeInterval: 60, IntervalsInHour: 1, IntervalsInDay: 24, IntervalsInWeek: 168},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("should return the appropriate config for a time interval of %d", tc.TimeInterval)
		t.Run(name, func(t *testing.T) {
			bTimeConfig, err := BuildConfigFromTimeInterval(tc.TimeInterval)

			if err != nil {
				t.Fatalf("expected no error, instead received %s", err.Error())
			}

			bStringSplitRegex, _ := regexp.Compile(fmt.Sprintf("(.{1,%d})", tc.IntervalsInHour))
			bStringDaySplitRegex, _ := regexp.Compile(fmt.Sprintf("(.{1,%d})", tc.IntervalsInDay))
			emptyHour := strings.Repeat(constants.ZeroPad, tc.IntervalsInHour)
			emptyDay := strings.Repeat(emptyHour, constants.HoursInDay)

			expected := &BTimeConfig{
				TimeInterval:            tc.TimeInterval,
				IntervalsInHour:         tc.IntervalsInHour,
				IntervalsInDay:          tc.IntervalsInDay,
				IntervalsInWeek:         tc.IntervalsInWeek,
				BStringSplitRegexStr:    bStringSplitRegex,
				BStringDaySplitRegexStr: bStringDaySplitRegex,
				EmptyHour:               emptyHour,
				EmptyDay:                emptyDay,
			}

			equal := expected.TimeInterval == bTimeConfig.TimeInterval
			if !equal {
				t.Fatalf("expected: %d, received: %d", expected.TimeInterval, bTimeConfig.TimeInterval)
			}
			equal = expected.IntervalsInHour == bTimeConfig.IntervalsInHour
			if !equal {
				t.Fatalf("expected: %d, received: %d", expected.IntervalsInHour, bTimeConfig.IntervalsInHour)
			}
			equal = expected.IntervalsInDay == bTimeConfig.IntervalsInDay
			if !equal {
				t.Fatalf("expected: %d, received: %d", expected.IntervalsInDay, bTimeConfig.IntervalsInDay)
			}
			equal = expected.IntervalsInWeek == bTimeConfig.IntervalsInWeek
			if !equal {
				t.Fatalf("expected: %d, received: %d", expected.IntervalsInWeek, bTimeConfig.IntervalsInWeek)
			}
			equal = expected.EmptyDay == bTimeConfig.EmptyDay
			if !equal {
				t.Fatalf("expected: %s, received: %s", expected.EmptyDay, bTimeConfig.EmptyDay)
			}
			equal = expected.EmptyHour == bTimeConfig.EmptyHour
			if !equal {
				t.Fatalf("expected: %s, received: %s", expected.EmptyHour, bTimeConfig.EmptyHour)
			}
		})
	}
}
