package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/keldonia/btime.go/models"
	"github.com/stretchr/testify/assert"
)

func TestNewBTimeFactoryValidInput(t *testing.T) {
	timeInterval := 5

	bTimeFactory, err := NewBTimeFactory(timeInterval)

	if err != nil {
		t.Fatalf("expected no error, received: %s", err.Error())
	}

	if bTimeFactory == nil {
		t.Fatalf("expected bTimeFactory to be generated")
	}
}

func TestNewBTimeFactoryinvalidInput(t *testing.T) {
	timeInterval := -1
	expectedErrorStr := fmt.Sprintf("[BConfig] received an invalid time interval: %d", timeInterval)

	bTimeFactory, err := NewBTimeFactory(timeInterval)

	if err.Error() != expectedErrorStr {
		t.Fatalf("expected error: %s, received: %s", expectedErrorStr, err.Error())
	}

	if bTimeFactory != nil {
		t.Fatalf("expected bTimeFactory to not be generated")
	}
}

func TestParseeBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := "01"
	returnInt := int64(1)

	bStringUtil.On("ParseBString", testArg).Return(&returnInt, nil)

	output, err := bTimeFactory.ParseBString(testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnInt, output)
}

func TestParseeBStringInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := "01"
	errorStr := "New Error"

	bStringUtil.On("ParseBString", testArg).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.ParseBString(testArg)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %d", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestGenerateBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	returnStr := "01"

	bStringUtil.On("GenerateBString", &testArg).Return(&returnStr, nil)

	output, err := bTimeFactory.GenerateBString(&testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnStr, output)
}

func TestGenerateBStringInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	errorStr := "New Error"

	bStringUtil.On("GenerateBString", &testArg).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.GenerateBString(&testArg)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %s", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestGenerateBStringFromAppointmentsProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := []models.Appointment{generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)}
	returnStrs := []string{"01"}

	bStringUtil.On("GenerateBStringFromAppointments", &testArg).Return(&returnStrs, nil)

	output, err := bTimeFactory.GenerateBStringFromAppointments(&testArg)

	assert.Nil(t, err)
	assert.Equal(t, &returnStrs, output)
}

func TestGenerateBStringFromAppointmentsInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := []models.Appointment{generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)}
	errorStr := "New Error"

	bStringUtil.On("GenerateBStringFromAppointments", &testArg).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.GenerateBStringFromAppointments(&testArg)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %s", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestTimeStringSplitProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
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
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg := float64(1)
	returnStrs := "01"

	bStringUtil.On("DecimalToBString", testArg).Return(returnStrs)

	output := bTimeFactory.DecimalToBString(testArg)

	assert.Equal(t, returnStrs, output)
}

func TestTestViabilityAndComputeInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := int64(1)
	testArg2 := int64(2)
	errorStr := "New Error"

	bScheduleUtil.On("TestViabilityAndCompute", testArg1, testArg2).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.TestViabilityAndCompute(testArg1, testArg2)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %d", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}
func TestTestViabilityAndComputeProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := int64(1)
	testArg2 := int64(2)
	returnArg := int64(3)

	bScheduleUtil.On("TestViabilityAndCompute", testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.TestViabilityAndCompute(testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestDeleteAppointmentInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	testArg2 := "00"
	errorStr := "New Error"

	bScheduleUtil.On("DeleteAppointment", &testArg1, testArg2).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.DeleteAppointment(&testArg1, testArg2)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %s", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestDeleteAppointmentProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := generateApptFromHoursAndMins(1, 10, 2, 20, 2, 2)
	testArg2 := "00"
	returnArg := "001"

	bScheduleUtil.On("DeleteAppointment", &testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.DeleteAppointment(&testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestDeleteAppointmentBStringInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := "01"
	testArg2 := "01"
	errorStr := "New Error"

	bScheduleUtil.On("DeleteAppointmentBString", testArg1, testArg2).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.DeleteAppointmentBString(testArg1, testArg2)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %s", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestDeleteAppointmentBStringProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := "01"
	testArg2 := "01"
	returnArg := "00"

	bScheduleUtil.On("DeleteAppointmentBString", testArg1, testArg2).Return(&returnArg, nil)

	output, err := bTimeFactory.DeleteAppointmentBString(testArg1, testArg2)

	assert.Nil(t, err)
	assert.Equal(t, &returnArg, output)
}

func TestModifyScheduleAndBookingInvalid(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
	}

	testArg1 := "001"
	testArg2 := "111"
	testArg3 := "110"
	errorStr := "New Error"

	bScheduleUtil.On("ModifyScheduleAndBooking", testArg1, testArg2, testArg3).Return(nil, fmt.Errorf(errorStr))

	output, err := bTimeFactory.ModifyScheduleAndBooking(testArg1, testArg2, testArg3)

	if output != nil {
		t.Fatalf("expected output to be nil, received: %s", *output)
	}

	if err.Error() != errorStr {
		t.Fatalf("expected error, received: %s", err.Error())
	}
}

func TestModifyScheduleAndBookingProperlyCalled(t *testing.T) {
	bTimeConfig, _ := BuildConfigFromTimeInterval(5)
	bStringUtil := NewMockBStringUtil(t)
	bScheduleUtil := NewMockBScheduleUtil(t)
	bConversionUtil := NewMockBConversionUtil(t)
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
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
	bPointerCalculator := NewMockBPointerCalculator(t)

	bTimeFactory := &BTimeFactoryImpl{
		bTimeConfig:        bTimeConfig,
		bStringUtil:        bStringUtil,
		bScheduleUtil:      bScheduleUtil,
		bConversionUtil:    bConversionUtil,
		bPointerCalculator: bPointerCalculator,
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

func TestFindWeekDay(t *testing.T) {
	timeInterval := 5

	bTimeFactory, _ := NewBTimeFactory(timeInterval)

	type test struct {
		Date    string
		WeekDay int
		Name    string
	}

	tests := []test{
		{
			Date:    "2020-02-02T00:00:00Z",
			WeekDay: 0,
			Name:    "Sunday",
		},
		{
			Date:    "2020-02-03T00:00:00Z",
			WeekDay: 1,
			Name:    "Monday",
		},
		{
			Date:    "2020-02-04T00:00:00Z",
			WeekDay: 2,
			Name:    "Tuesday",
		},
		{
			Date:    "2020-02-05T00:00:00Z",
			WeekDay: 3,
			Name:    "Wednesday",
		},
		{
			Date:    "2020-02-06T00:00:00Z",
			WeekDay: 4,
			Name:    "Thursday",
		},
		{
			Date:    "2020-02-07T00:00:00Z",
			WeekDay: 5,
			Name:    "Friday",
		},
		{
			Date:    "2020-02-08T00:00:00Z",
			WeekDay: 6,
			Name:    "Saturday",
		},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		t.Run(tc.Name, func(t *testing.T) {
			baseDate, _ := time.Parse("2006-01-02T15:04:05Z", tc.Date)
			weekday := bTimeFactory.FindWeekDay(&baseDate)

			if weekday != tc.WeekDay {
				t.Fatalf("expected weekday int: %d, received: %d for weekday: %s", tc.WeekDay, weekday, tc.Name)
			}
		})
	}
}
