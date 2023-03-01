package models

import "time"

type Appointment struct {
	StartTime time.Time
	EndTime   time.Time
}

type AppointmentDuo struct {
	InitialAppointment time.Time
	SecondAppointment  time.Time
}
