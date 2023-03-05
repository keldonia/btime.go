package core

import (
	"testing"
	"time"

	"github.com/keldonia/btime.go/models"
	"github.com/stretchr/testify/assert"
)

func TestParseeBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg := "01"
	returnInt := int64(1)

	bStringUtil.On("ParseBString", testArg).Return(&returnInt, nil)

	output, err := bTimeFactory.ParseBString(testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnInt, output)
}

func TestGenerateBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	returnStr := "01"

	bStringUtil.On("GenerateBString", &testArg).Return(&returnStr, nil)

	output, err := bTimeFactory.GenerateBString(&testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnStr, output)
}

func TestGenerateBStringFromAppointmentsProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg := []models.Appointment{generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)}
	returnStrs := []string{"01"}

	bStringUtil.On("GenerateBStringFromAppointments", &testArg).Return(&returnStrs, nil)

	output, err := bTimeFactory.GenerateBStringFromAppointments(&testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnStrs, output)
}

func TestTimeStringSplitProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg := bTimeConfig.EmptyDay
	returnStrs := []string{"00"}

	bStringUtil.On("TimeStringSplit", testArg).Return(returnStrs, nil)

	output := bTimeFactory.TimeStringSplit(testArg)

	assert.Equal(t, returnStrs, output)
}

func TestDecimalToBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg := float64(1)
	returnStrs := "01"

	bStringUtil.On("DecimalToBString", testArg).Return(returnStrs)

	output := bTimeFactory.DecimalToBString(testArg)

	assert.Equal(t, returnStrs, output)
}

func TestTestViabilityAndComputeProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg1 := int64(1)
	testArg2 := int64(2)
	returnArg := int64(3)

	bScheduleUtil.On("TestViabilityAndCompute", testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.TestViabilityAndCompute(testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestDeleteAppointmentProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg1 := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	testArg2 := "00"
	returnArg := "001"

	bScheduleUtil.On("DeleteAppointment", &testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.DeleteAppointment(&testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestDeleteAppointmentBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg1 := "01"
	testArg2 := "01"
	returnArg := "00"

	bScheduleUtil.On("DeleteAppointmentBString", testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.DeleteAppointmentBString(testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestModifyScheduleAndBookingProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	testArg1 := "001"
	testArg2 := "111"
	testArg3 := "110"
	returnArg := "000"

	bScheduleUtil.On("ModifyScheduleAndBooking", testArg1, testArg2, testArg3).Return(&returnArg, nil)

	output, err := bTimeFactory.ModifyScheduleAndBooking(testArg1, testArg2, testArg3)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestConvertScheduleToAppointmentScheduleProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:     bTimeConfig,
		bStringUtil:     bStringUtil,
		bScheduleUtil:   bScheduleUtil,
		bConversionUtil: bConversionUtil,
	}

	baseDate, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-09T00:00:00Z")
	emptySchedule := generateEmptyWeek(5)
	emptyBookings := generateEmptyWeek(5)
	schedule := models.Schedule{
		Schedule:  &emptySchedule,
		Bookings:  &emptyBookings,
		WeekStart: &baseDate,
	}
	emptyAvail := generateEmptyWeek(5)

	returnArg := models.AppointmentSchedule{
		Schedule:     &[][]models.Appointment{},
		Bookings:     &[][]models.Appointment{},
		Availability: &[][]models.Appointment{},
		WeekStart:    &baseDate,
	}

	bConversionUtil.On("ConvertScheduleToAppointmentSchedule", &schedule, emptyAvail).Return(&returnArg)

	output := bTimeFactory.ConvertScheduleToAppointmentSchedule(&schedule, emptyAvail)

	assert.Equal(t, &returnArg, output)
}
