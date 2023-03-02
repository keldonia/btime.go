package core

import (
	"fmt"
	"math"
	"time"

	"github.com/keldonia/btime.go/constants"
)

type BPointerCalculator struct {
	bTimeConfig *BTimeConfig
}

func NewBPointerCalculator(bTimeConfig *BTimeConfig) (*BConversionUtil, error) {
	if bTimeConfig == nil {
		return nil, fmt.Errorf("No BTimeConfig was provided")
	}

	return &BConversionUtil{
		bTimeConfig: bTimeConfig,
	}, nil
}

func (bpc *BPointerCalculator) FindBPointerIncludingDay(date *time.Time) int {
	hourAndMinutePointer := bpc.FindBPointer(date)
	dayModifier := bpc.FindBPointerModiferForDayOfWeek(date)

	return dayModifier + hourAndMinutePointer
}

func (bpc *BPointerCalculator) FindBPointerModiferForDayOfWeek(date *time.Time) int {
	return date.UTC().Day() * bpc.bTimeConfig.IntervalsInHour * constants.HoursInDay
}

func (bpc *BPointerCalculator) FindBPointer(date *time.Time) int {
	hourPointer := date.UTC().Hour() * bpc.bTimeConfig.IntervalsInHour
	minutePointer := int(math.Floor(float64(date.UTC().Minute()) / float64(bpc.bTimeConfig.TimeInterval)))

	return hourPointer + minutePointer
}
