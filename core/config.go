package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/keldonia/btime.go/constants"
	"golang.org/x/exp/slices"
)

type BTimeConfig struct {
	TimeInterval            int
	IntervalsInHour         int
	IntervalsInDay          int
	IntervalsInWeek         int
	BStringSplitRegexStr    *regexp.Regexp
	BStringDaySplitRegexStr *regexp.Regexp
	EmptyHour               string
	EmptyDay                string
}

func BuildConfigFromTimeInterval(timeInterval int) (*BTimeConfig, error) {
	validTimeIntervals := constants.ValidTimeIntervals()

	if !slices.Contains(validTimeIntervals, timeInterval) {
		return nil, fmt.Errorf("[BConfig] received an invalid time interval: %d", timeInterval)
	}

	intervalsInHour := constants.MinutesInHour / timeInterval
	intervalsInDay := intervalsInHour * constants.HoursInDay
	intervalsInWeek := intervalsInDay * constants.DaysInWeek
	bStringSplitRegexStr := fmt.Sprintf("(.{1,%d})", intervalsInHour)
	bStringSplitRegex, _ := regexp.Compile(bStringSplitRegexStr)

	bStringDaySplitRegexStr := fmt.Sprintf("(.{1,%d})", intervalsInDay)
	// We have already checked that the intervalsInDay will be valid
	bStringDaySplitRegex, _ := regexp.Compile(bStringDaySplitRegexStr)

	emptyHour := strings.Repeat(constants.ZeroPad, intervalsInHour)
	emptyDay := strings.Repeat(emptyHour, constants.HoursInDay)

	return &BTimeConfig{
		TimeInterval:            timeInterval,
		IntervalsInHour:         intervalsInHour,
		IntervalsInDay:          intervalsInDay,
		IntervalsInWeek:         intervalsInWeek,
		BStringSplitRegexStr:    bStringSplitRegex,
		BStringDaySplitRegexStr: bStringDaySplitRegex,
		EmptyHour:               emptyHour,
		EmptyDay:                emptyDay,
	}, nil
}
