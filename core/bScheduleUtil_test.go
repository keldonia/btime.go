package core

import (
	"fmt"
	"strings"
	"testing"

	"github.com/keldonia/btime.go/models"
)

func TestMergeScheduleBStringWithTest(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Appt1    string
		Appt2    string
		Error    bool
		Expected string
	}

	tests := []test{
		{Appt1: "000011110000", Appt2: "000000000011", Error: false, Expected: "000011110011"},
		{Appt1: "000000000000", Appt2: "000000000011", Error: false, Expected: "000000000011"},
		{Appt1: "000011110000", Appt2: "000000000000", Error: false, Expected: "000011110000"},
		{Appt1: "011000000000", Appt2: "000000011000", Error: false, Expected: "011000011000"},
		{Appt1: "100000000000", Appt2: "000000011111", Error: false, Expected: "100000011111"},
		{Appt1: "011110000000", Appt2: "000011110000", Error: true, Expected: ""},
		{Appt1: "110000000000", Appt2: "111100000000", Error: true, Expected: ""},
		{Appt1: "000000000111", Appt2: "000000111110", Error: true, Expected: ""},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("should expect binary appts of %s & %s to return %s", tc.Appt1, tc.Appt2, tc.Expected)
		t.Run(name, func(t *testing.T) {
			computed, err := bScheduleUtil.MergeScheduleBStringWithTest(tc.Appt1, tc.Appt2)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *computed) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *computed)
			}
		})
	}
}

func TestLoopMergeScheduleBStringWithTest(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Appt1    models.Appointment
		Appt2    models.Appointment
		Error    bool
		Expected string
	}

	two := 2

	tests := []test{
		{Appt1: generateApptFromHoursAndMins(1, 0, 1, 24, two, two), Appt2: generateApptFromHoursAndMins(0, 12, 0, 24, two, two), Error: false, Expected: "001110000000111110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{Appt1: generateApptFromHoursAndMins(1, 0, 1, 24, two, two), Appt2: generateApptFromHoursAndMins(4, 12, 5, 24, two, two), Error: false, Expected: "000000000000111110000000000000000000000000000000001111111111111110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 40, two, two), Appt2: generateApptFromHoursAndMins(0, 12, 0, 24, two, two), Error: true, Expected: ""},
		{Appt1: generateApptFromHoursAndMins(12, 20, 13, 40, two, two), Appt2: generateApptFromHoursAndMins(13, 12, 15, 24, two, two), Error: true, Expected: ""},
		{Appt1: generateApptFromHoursAndMins(13, 20, 12, 40, two, two), Appt2: generateApptFromHoursAndMins(13, 12, 15, 24, two, two), Error: true, Expected: ""},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 20, two, two), Appt2: generateApptFromHoursAndMins(0, 20, 0, 24, two, two), Error: true, Expected: ""},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			bString, _ := bScheduleUtil.bStringUtil.GenerateBString(&tc.Appt2)

			mergedBString, err := bScheduleUtil.MergeScheduleBStringsWithTest(&tc.Appt1, *bString)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *mergedBString) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *mergedBString)
			}
		})
	}
}

func TestModifyScheduleAndBooking(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Base     string
		Test     string
		Appt     string
		Error    bool
		Expected string
	}

	tests := []test{
		{Base: "000011110000", Test: "000011111111", Appt: "000000000011", Error: false, Expected: "000011110011"},
		{Base: "000000000000", Test: "000000000000", Appt: "000000000011", Error: true, Expected: ""},
		{Base: "000011110000", Test: "000011110000", Appt: "000000000000", Error: false, Expected: "000011110000"},
		{Base: "011000000000", Test: "000000000000", Appt: "000000011000", Error: true, Expected: ""},
		{Base: "100000000000", Test: "111111111111", Appt: "000000011111", Error: false, Expected: "100000011111"},
		{Base: "011110000000", Test: "011110011111", Appt: "000000011111", Error: false, Expected: "011110011111"},
		{Base: "011110000000", Test: "000000000000", Appt: "000011110000", Error: true, Expected: ""},
		{Base: "110000000000", Test: "000000000000", Appt: "111100000000", Error: true, Expected: ""},
		{Base: "000000000111", Test: "000000000000", Appt: "000000111110", Error: true, Expected: ""},
		{Base: "000000000111", Test: "111111111111", Appt: "000000111110", Error: true, Expected: ""},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			mergedBString, err := bScheduleUtil.ModifyScheduleAndBooking(tc.Base, tc.Test, tc.Appt)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *mergedBString) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *mergedBString)
			}
		})
	}
}

func TestModifyScheduleAndBookingMultipleIntervals(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	scheduleToModify := generateApptFromHoursAndMins(12, 40, 17, 40, 2, 2)
	baseAvail := generateApptFromHoursAndMins(12, 20, 19, 24, 2, 2)
	appt := generateApptFromHoursAndMins(12, 20, 12, 39, 2, 2)
	expected := generateApptFromHoursAndMins(12, 20, 17, 40, 2, 2)

	scheduleToModifyStr, _ := bStringUtil.GenerateBString(&scheduleToModify)
	baseAvailStr, _ := bStringUtil.GenerateBString(&baseAvail)
	apptStr, _ := bStringUtil.GenerateBString(&appt)
	expectedStr, _ := bStringUtil.GenerateBString(&expected)

	computed, _ := bScheduleUtil.ModifyScheduleAndBooking(*scheduleToModifyStr, *baseAvailStr, *apptStr)

	if strings.Compare(*expectedStr, *computed) != 0 {
		t.Fatalf("expected %s, received %s", *expectedStr, *computed)
	}
}

func TestModifyScheduleAndBookingInterval(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Base     string
		Test     string
		Appt     string
		Error    bool
		Expected string
	}

	tests := []test{
		{Base: "000011110000", Test: "000011111111", Appt: "000000000011", Error: false, Expected: "000011110011"},
		{Base: "000000000000", Test: "000000000000", Appt: "000000000011", Error: true, Expected: ""},
		{Base: "000011110000", Test: "000011110000", Appt: "000000000000", Error: false, Expected: "000011110000"},
		{Base: "011000000000", Test: "000000000000", Appt: "000000011000", Error: true, Expected: ""},
		{Base: "100000000000", Test: "111111111111", Appt: "000000011111", Error: false, Expected: "100000011111"},
		{Base: "011110000000", Test: "011110011111", Appt: "000000011111", Error: false, Expected: "011110011111"},
		{Base: "011110000000", Test: "000000000000", Appt: "000011110000", Error: true, Expected: ""},
		{Base: "110000000000", Test: "000000000000", Appt: "111100000000", Error: true, Expected: ""},
		{Base: "000000000111", Test: "000000000000", Appt: "000000111110", Error: true, Expected: ""},
		{Base: "000000000111", Test: "111111111111", Appt: "000000111110", Error: true, Expected: ""},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			mergedBString, err := bScheduleUtil.ModifyScheduleAndBookingInterval(tc.Base, tc.Test, tc.Appt)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *mergedBString) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *mergedBString)
			}
		})
	}
}

func TestDeleteAppointment(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Appt1    models.Appointment
		Appt2    models.Appointment
		Error    bool
		Expected string
	}

	tests := []test{
		{Appt1: generateApptFromHoursAndMins(1, 0, 1, 24, 2, 2), Appt2: generateApptFromHoursAndMins(0, 12, 1, 24, 2, 2), Error: false, Expected: "001111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{Appt1: generateApptFromHoursAndMins(1, 0, 1, 24, 2, 2), Appt2: generateApptFromHoursAndMins(4, 12, 5, 24, 2, 2), Error: true, Expected: "BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 111110000000 Schedule: 000000000000"},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 40, 2, 2), Appt2: generateApptFromHoursAndMins(0, 12, 0, 40, 2, 2), Error: false, Expected: "001100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{Appt1: generateApptFromHoursAndMins(12, 20, 13, 40, 2, 2), Appt2: generateApptFromHoursAndMins(13, 12, 15, 24, 2, 2), Error: true, Expected: "BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 000011111111 Schedule: 000000000000"},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 40, 2, 2), Appt2: generateApptFromHoursAndMins(0, 12, 12, 40, 2, 2), Error: false, Expected: "001100000111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 40, 2, 2), Appt2: generateApptFromHoursAndMins(0, 12, 0, 24, 2, 2), Error: true, Expected: "BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 000011111000 Schedule: 001110000000"},
		{Appt1: generateApptFromHoursAndMins(0, 20, 0, 20, 2, 2), Appt2: generateApptFromHoursAndMins(0, 20, 0, 24, 2, 2), Error: false, Expected: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			appt2Str, _ := bStringUtil.GenerateBString(&tc.Appt2)
			mergedBString, err := bScheduleUtil.DeleteAppointment(&tc.Appt1, *appt2Str)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *mergedBString) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *mergedBString)
			}
		})
	}
}

func TestDeleteAppointmentShouldFailIfPassedInvalidApptToDelete(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	invalidAppt := generateApptFromHoursAndMins(12, 20, 11, 40, 2, 2)
	appt2 := generateApptFromHoursAndMins(13, 12, 15, 24, 2, 2)

	appt2Str, _ := bStringUtil.GenerateBString(&appt2)

	_, err := bScheduleUtil.DeleteAppointment(&invalidAppt, *appt2Str)

	if err == nil {
		t.Fatal("Did not received expected error")
	}
}

func TestDeleteAppointmentInterval(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Base     string
		Appt     string
		Error    bool
		Expected string
	}

	tests := []test{
		{Base: "000011110011", Appt: "000000000011", Error: false, Expected: "000011110000"},
		{Base: "000000000000", Appt: "000000000011", Error: true, Expected: `BScheduleUtil Error: invalid deletion, interval to delete occurs outside of schedule interval. To be deleted: 000000000011 Schedule: 000000000000`},
		{Base: "000011110000", Appt: "000000000000", Error: false, Expected: "000011110000"},
		{Base: "011000011000", Appt: "000000011000", Error: false, Expected: "011000000000"},
		{Base: "100000011111", Appt: "000000011111", Error: false, Expected: "100000000000"},
		{Base: "011110000000", Appt: "000110000000", Error: false, Expected: "011000000000"},
		{Base: "111100000000", Appt: "110000000000", Error: false, Expected: "001100000000"},
		{Base: "000000111110", Appt: "000000000110", Error: false, Expected: "000000111000"},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			mergedBString, err := bScheduleUtil.DeleteAppointmentInterval(tc.Appt, tc.Base)

			if tc.Error && err == nil {
				t.Fatalf("did not received expected error")
			}

			if !tc.Error && strings.Compare(tc.Expected, *mergedBString) != 0 {
				t.Fatalf("expected %s, received %s", tc.Expected, *mergedBString)
			}
		})
	}
}

func TestValidDeletion(t *testing.T) {
	timeInterval := 5
	bTimeConfig, _ := BuildConfigFromTimeInterval(timeInterval)
	bStringUtil, _ := NewBStringUtil(bTimeConfig)
	bScheduleUtil, _ := NewBScheduleUtil(bTimeConfig)

	type test struct {
		Base     string
		Appt     string
		Expected bool
	}

	tests := []test{
		{Base: "000011110011", Appt: "000000000011", Expected: true},
		{Base: "000000000000", Appt: "000000000011", Expected: false},
		{Base: "000011110000", Appt: "000000000000", Expected: true},
		{Base: "011000011000", Appt: "000000011000", Expected: true},
		{Base: "100000011111", Appt: "000000011111", Expected: true},
		{Base: "011110000000", Appt: "000110000000", Expected: true},
		{Base: "111100000000", Appt: "110000000000", Expected: true},
		{Base: "000000111110", Appt: "000000000110", Expected: true},
	}

	for i := 0; i < len(tests); i++ {
		tc := tests[i]
		name := fmt.Sprintf("tc: %d", i)
		t.Run(name, func(t *testing.T) {
			parsedBase, _ := bStringUtil.ParseBString(tc.Base)
			parsedAppt, _ := bStringUtil.ParseBString(tc.Appt)

			validDeletion := bScheduleUtil.validDeletion(*parsedBase, *parsedAppt)

			if tc.Expected != validDeletion {
				t.Fatalf("expected %t, received %t", tc.Expected, validDeletion)
			}
		})
	}
}
