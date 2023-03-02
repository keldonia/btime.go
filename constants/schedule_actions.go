package constants

type ScheduleAction string

const (
	DELETE_APPOINTMENT ScheduleAction = "DELETE_APPOINTMENT"
	BOOKING_UPDATE     ScheduleAction = "BOOKING_UPDATE"
	UNKNOWN            ScheduleAction = "UNKNOWN"
)
