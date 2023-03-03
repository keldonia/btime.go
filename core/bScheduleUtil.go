package core

import (
	"fmt"
	"strings"

	"github.com/keldonia/btime.go/constants"
	"github.com/keldonia/btime.go/models"
)

type BScheduleUtil struct {
	bTimeConfig *BTimeConfig
	bStringUtil *BStringUtil
}

func NewBScheduleUtil(bTimeConfig *BTimeConfig) (*BScheduleUtil, error) {
	if bTimeConfig == nil {
		return nil, fmt.Errorf("[BScheduleUtil] No BTimeConfig was provided")
	}

	bStringUtil, err := NewBStringUtil(bTimeConfig)

	if err != nil {
		return nil, err
	}

	return &BScheduleUtil{
		bTimeConfig: bTimeConfig,
		bStringUtil: bStringUtil,
	}, nil
}

func (bschu *BScheduleUtil) MergeScheduleBStringsWithTest(timeSlot *models.Appointment, schedule string) (*string, error) {
	if timeSlot.EndTime.Before(*timeSlot.StartTime) {
		return nil, fmt.Errorf("BSchedule Error: Invalid timeslot passed to merge schedule BString StartTime: %s EndTime: %s", timeSlot.StartTime.UTC().GoString(), timeSlot.EndTime.UTC().GoString())
	}
	apptBString, err := bschu.bStringUtil.GenerateBString(timeSlot)

	if err != nil {
		return nil, err
	}

	mergedBString, err := bschu.MergeScheduleBStringsWithTestBase(*apptBString, schedule)

	if err != nil {
		return nil, err
	}

	return mergedBString, nil
}

func (bschu *BScheduleUtil) MergeScheduleBStringsWithTestBase(apptBString string, schedule string) (*string, error) {
	apptMask := bschu.bStringUtil.TimeStringSplit(apptBString)
	splitSchedule := bschu.bStringUtil.TimeStringSplit(schedule)
	mergedSchedule := []string{}

	//  NB: Iterate over each section of the schedule & appt to
	//  generate a combined schedule, it returns early if the merged
	//  schedule and appt BStrings conflict
	for i := 0; i < len(splitSchedule); i++ {
		mergeReturn, err := bschu.MergeScheduleBStringWithTest(splitSchedule[i], apptMask[i])

		if err != nil {
			return nil, err
		}
		mergedSchedule = append(mergedSchedule, *mergeReturn)
	}

	joinedSchedule := strings.Join(mergedSchedule, "")

	return &joinedSchedule, nil
}

func (bschu *BScheduleUtil) MergeScheduleBStringWithTest(timeSlotBString string, schedule string) (*string, error) {
	parsedSchedule, err := bschu.bStringUtil.ParseBString(schedule)

	if err != nil {
		return nil, err
	}

	parsedApptBString, err := bschu.bStringUtil.ParseBString(timeSlotBString)

	if err != nil {
		return nil, err
	}

	// Performs a XOR on the schedule and the proposed schedule
	modified := *parsedSchedule ^ *parsedApptBString
	// Performs an OR on the schedule and the proposed schedule
	test := *parsedSchedule | *parsedApptBString

	if modified != test {
		return nil, fmt.Errorf("BScheduleUtil Error: Schedules conflict and overlap")
	}

	modifiedBString := bschu.bStringUtil.DecimalToBString(float64(modified))

	return &modifiedBString, nil
}

func (bschu *BScheduleUtil) ModifyScheduleAndBooking(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	splitToModify := bschu.bStringUtil.TimeStringSplit(scheduleBStringToModify)
	splitToTest := bschu.bStringUtil.TimeStringSplit(scheduleBStringToTest)
	splitAppt := bschu.bStringUtil.TimeStringSplit(appt)
	modifiedSchedule := []string{}

	//  NB: Iterate over each section of the schedule & appt to
	//  generate a combined schedule, it returns early if the merged
	//  schedule and appt BStrings conflict
	for i := 0; i < len(splitToModify); i++ {
		mergeReturn, err := bschu.ModifyScheduleAndBookingInterval(splitToModify[i], splitToTest[i], splitAppt[i])

		if err != nil {
			return nil, err
		}

		modifiedSchedule = append(modifiedSchedule, *mergeReturn)
	}

	joinedSchedule := strings.Join(modifiedSchedule, "")

	return &joinedSchedule, nil
}

func (bschu *BScheduleUtil) ModifyScheduleAndBookingInterval(scheduleBStringToModify string, scheduleBStringToTest string, appt string) (*string, error) {
	parsedToModify, err := bschu.bStringUtil.ParseBString(scheduleBStringToModify)

	if err != nil {
		return nil, err
	}
	// Flip the bits to test the pattern
	parsedToTest, err := bschu.bStringUtil.ParseBString(scheduleBStringToTest)

	if err != nil {
		return nil, err
	}

	parsedApptBString, err := bschu.bStringUtil.ParseBString(appt)

	if err != nil {
		return nil, err
	}

	_, err = bschu.TestViabilityAndCompute(*parsedApptBString, *parsedToTest)

	if err != nil {
		return nil, err
	}

	update, err := bschu.TestViabilityAndCompute(*parsedApptBString, *parsedToModify)

	if err != nil {
		return nil, err
	}

	updateStr := bschu.bStringUtil.DecimalToBString(float64(*update))

	return &updateStr, nil
}

func (bschu *BScheduleUtil) TestViabilityAndCompute(binary1 int64, binary2 int64) (*int64, error) {
	modified := binary1 ^ binary2
	test := binary1 | binary2

	if modified == test {
		return &modified, nil
	}

	return nil, fmt.Errorf("BScheduleUtil Error: Time intervals overlap.")
}

func (bschu *BScheduleUtil) DeleteAppointment(timeSlotToDelete *models.Appointment, scheduleSlot string) (*string, error) {
	if timeSlotToDelete.EndTime.Before(*timeSlotToDelete.StartTime) {
		return nil, fmt.Errorf("BSchedule Error: Invalid appointment passed to delete appointment StartTime: %s EndTime: %s", timeSlotToDelete.StartTime.UTC().GoString(), timeSlotToDelete.EndTime.UTC().GoString())
	}

	apptToDeleteBString, err := bschu.bStringUtil.GenerateBString(timeSlotToDelete)

	if err != nil {
		return nil, err
	}

	deleteApptString, err := bschu.DeleteAppointmentBString(*apptToDeleteBString, scheduleSlot)

	if err != nil {
		return nil, err
	}

	return deleteApptString, nil
}

func (bschu *BScheduleUtil) DeleteAppointmentBString(bStringToDelete string, scheduleSlot string) (*string, error) {
	apptMask := bschu.bStringUtil.TimeStringSplit(bStringToDelete)
	splitBookings := bschu.bStringUtil.TimeStringSplit(scheduleSlot)
	returnAppointments := []string{}

	for i := 0; i < constants.HoursInDay; i++ {
		currentInterval, err := bschu.DeleteAppointmentInterval(apptMask[i], splitBookings[i])
		if err != nil {
			return nil, err
		}

		returnAppointments = append(returnAppointments, *currentInterval)
	}

	completeBString := strings.Join(returnAppointments, "")

	return &completeBString, nil
}

func (bschu *BScheduleUtil) DeleteAppointmentInterval(timeSlotBString string, scheduleInterval string) (*string, error) {
	parsedSchedule, err := bschu.bStringUtil.ParseBString(scheduleInterval)

	if err != nil {
		return nil, err
	}

	parsedApptBString, err := bschu.bStringUtil.ParseBString(timeSlotBString)
	if err != nil {
		return nil, err
	}

	if !bschu.validDeletion(*parsedSchedule, *parsedApptBString) {
		return nil, fmt.Errorf("BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: %s Schedule: %s", timeSlotBString, scheduleInterval)
	}

	// Performs a XOR on the schedule and the proposed schedule
	modified := *parsedSchedule ^ *parsedApptBString
	modifiedBString := bschu.bStringUtil.DecimalToBString(float64(modified))

	return &modifiedBString, nil
}

func (bschu *BScheduleUtil) validDeletion(baseNumber int64, toDeleteNumber int64) bool {
	orTest := baseNumber | toDeleteNumber

	return orTest == baseNumber
}
