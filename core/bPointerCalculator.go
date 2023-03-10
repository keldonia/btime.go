package core

import (
	"fmt"
	"math"
	"time"

	"github.com/keldonia/btime.go/constants"
)

//go:generate mockery --name BPointerCalculatorImpl
type BPointerCalculator interface {
	FindBPointerIncludingDay(date *time.Time) int
	FindBPointerModiferForDayOfWeek(date *time.Time) int
	FindBPointer(date *time.Time) int
}

type BPointerCalculatorImpl struct {
	bTimeConfig *BTimeConfig
}

// Instantiates a new BPointerCalculatorImpl, which is responsible for calculating bPointers, for use by other bTime classes
func NewBPointerCalculator(bTimeConfig *BTimeConfig) (BPointerCalculator, error) {
	if bTimeConfig == nil {
		return nil, fmt.Errorf("[BPointerCalculator] No BTimeConfig was provided")
	}

	return &BPointerCalculatorImpl{
		bTimeConfig: bTimeConfig,
	}, nil
}

// Finds a the pointer for a given date in time based on the instatiated time interval, including day of the week
func (bpc *BPointerCalculatorImpl) FindBPointerIncludingDay(date *time.Time) int {
	hourAndMinutePointer := bpc.FindBPointer(date)
	dayModifier := bpc.FindBPointerModiferForDayOfWeek(date)

	return dayModifier + hourAndMinutePointer
}

// Finds the pointer modifer to correct for day of the week
func (bpc *BPointerCalculatorImpl) FindBPointerModiferForDayOfWeek(date *time.Time) int {
	weekday := date.UTC().Weekday()

	// Shift pointer to start week on Sunday
	if weekday == time.Sunday {
		weekday = 0
	}

	return int(weekday) * bpc.bTimeConfig.IntervalsInHour * constants.HoursInDay
}

// Finds a the pointer for a given date in time based on the instatiated time interval within a given day
func (bpc *BPointerCalculatorImpl) FindBPointer(date *time.Time) int {
	hourPointer := date.UTC().Hour() * bpc.bTimeConfig.IntervalsInHour
	minutePointer := int(math.Floor(float64(date.UTC().Minute()) / float64(bpc.bTimeConfig.TimeInterval)))

	return hourPointer + minutePointer
}
