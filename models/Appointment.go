package models

import "time"

type Appointment struct {
	StartTime *time.Time
	EndTime   *time.Time
}

type AppointmentDuo struct {
	InitialAppointment *Appointment
	SecondAppointment  *Appointment
}
