package models

import "time"

type Schedule struct {
	Schedule  []string // This indexes 0-6, starting (0) with Sunday
	Bookings  []string // This indexes 0-6, starting (0) with Sunday
	WeekStart time.Time
}

type AppointmentSchedule struct {
	Schedule [][]Appointment
}
