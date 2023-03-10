package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/keldonia/btime.go/constants"
	"github.com/keldonia/btime.go/models"
)

//go:generate mockery --name BStringUtil
type BStringUtil interface {
	GenerateBString(appt *models.Appointment) (*string, error)
	GenerateBStringFromAppointments(appointments *[]models.Appointment) (*[]string, error)
	TimeStringSplit(scheduleString string) []string
	ParseBString(bString string) (*int64, error)
	DecimalToBString(decimal float64) string
}

type BStringUtilImpl struct {
	bTimeConfig        *BTimeConfig
	bPointerCalculator BPointerCalculator
	emptyHour          string
	emptyDay           string
}

// Instantiates a new BStringUtil, which is responsible for
// generating and formatting  the bStrings used by this package
//
// NB: Typically a temporal resolution of 5 mins is sufficient,
// as it constitutes the smallest billable unit in most juristictions
func NewBStringUtil(bTimeConfig *BTimeConfig) (BStringUtil, error) {
	if bTimeConfig == nil {
		return nil, fmt.Errorf("[BStringUtil] No BTimeConfig was provided")
	}

	// We already checked that bTimeConfig will not be nil
	bPointerCalculator, _ := NewBPointerCalculator(bTimeConfig)

	emptyHour := strings.Repeat(constants.ZeroPad, bTimeConfig.IntervalsInHour)
	emptyDay := strings.Repeat(emptyHour, constants.HoursInDay)

	return &BStringUtilImpl{
		bTimeConfig:        bTimeConfig,
		bPointerCalculator: bPointerCalculator,
		emptyHour:          emptyHour,
		emptyDay:           emptyDay,
	}, nil
}

// Generates a bString representation of a given appointment, assuming it is valid.
// If the appointment is invalid, it will throw an error
func (bsu *BStringUtilImpl) GenerateBString(appt *models.Appointment) (*string, error) {
	if appt.EndTime.Before(*appt.StartTime) {
		return nil, fmt.Errorf("BString Error: Appointment can't end before it begins.  Appointment start: %s Appointment end: %s", appt.StartTime.UTC().GoString(), appt.EndTime.UTC().GoString())
	}

	startPointer := bsu.bPointerCalculator.FindBPointer(appt.StartTime)
	endPointer := bsu.bPointerCalculator.FindBPointer(appt.EndTime)
	timeBlock := endPointer - startPointer + 1

	startStr := bsu.emptyDay[0:startPointer]
	occupied := strings.Repeat(constants.OnePad, timeBlock)
	endStr := bsu.emptyDay[endPointer+1 : len(bsu.emptyDay)]

	bString := fmt.Sprintf("%s%s%s", startStr, occupied, endStr)

	return &bString, nil
}

// Generates a bString representation of a given array of appointments, assuming it is valid.
// If the appointment is invalid, it will throw an error
//
// NB: This method generates a representation of the entire week
//
// NB: Assumes appointments in array don't overlap
func (bsu *BStringUtilImpl) GenerateBStringFromAppointments(appointments *[]models.Appointment) (*[]string, error) {
	var composedBString string = ""

	for i := 0; i < len(*appointments); i++ {
		appt := (*appointments)[i]

		if appt.StartTime.After(*appt.EndTime) {
			return nil, fmt.Errorf("BString Error: Appointment can't end before it begins.  Appointment start: %s Appointment end: %s", appt.StartTime.UTC().GoString(), appt.EndTime.UTC().GoString())
		}

		startPointer := bsu.bPointerCalculator.FindBPointerIncludingDay(appt.StartTime)
		endPointer := bsu.bPointerCalculator.FindBPointerIncludingDay(appt.EndTime)
		timeBlock := endPointer - startPointer + 1

		// If an appt begins before the previous ends, it is invalid
		if startPointer < len(composedBString) {
			return nil, fmt.Errorf("BString Error: Appointment can't begin before previous appointment ends.  Appointment start: %s Previous Appointment end: %s", appt.StartTime.UTC().GoString(), appt.EndTime.UTC().GoString())
		}

		// Adds padding between appointments
		zeroesToAdd := startPointer - len(composedBString)
		addedZeros := strings.Repeat(constants.ZeroPad, zeroesToAdd)
		composedBString = composedBString + addedZeros + strings.Repeat(constants.OnePad, timeBlock)
	}

	// Pad out remainder of week
	endOfWeekPadding := bsu.bTimeConfig.IntervalsInWeek - len(composedBString)
	composedBString = composedBString + strings.Repeat(constants.ZeroPad, endOfWeekPadding)

	splitString := bsu.bTimeConfig.BStringDaySplitRegexStr.FindAllString(composedBString, -1)

	return &splitString, nil
}

// Splits each schedule BString into a string of length defined in the regex
func (bsu *BStringUtilImpl) TimeStringSplit(scheduleString string) []string {
	return bsu.bTimeConfig.BStringSplitRegexStr.FindAllString(scheduleString, -1)
}

// Converts bString representation of a number into a number for calculation purposes
func (bsu *BStringUtilImpl) ParseBString(bString string) (*int64, error) {
	numeric, err := strconv.ParseInt(bString, 2, 64)

	if err != nil {
		return nil, err
	}

	return &numeric, nil
}

// Converts number into a bString representation with the given scheduling interval
func (bsu *BStringUtilImpl) DecimalToBString(decimal float64) string {
	return fmt.Sprintf("%0*s", bsu.bTimeConfig.IntervalsInHour, strconv.FormatInt(int64(decimal), constants.BinaryBase))
}
