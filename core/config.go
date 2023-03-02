package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/keldonia/btime.go/constants"
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
	intervalsInHour := constants.MinutesInHour / timeInterval
	intervalsInDay := intervalsInHour * constants.HoursInDay
	intervalsInWeek := intervalsInDay * constants.DaysInWeek
	bStringSplitRegexStr := fmt.Sprintf("(.{1,%d})", intervalsInHour)
	bStringSplitRegex, err := regexp.Compile(bStringSplitRegexStr)

	if err != nil {
		return nil, err
	}

	bStringDaySplitRegexStr := fmt.Sprintf("(.{1,%d})", intervalsInDay)
	bStringDaySplitRegex, err := regexp.Compile(bStringDaySplitRegexStr)
	if err != nil {
		return nil, err
	}

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
