package main

import (
	"time"

	"github.com/keldonia/btime.go/src/models"
)

func main() {
	appt := models.Appointment{
		StartTime: time.Now(),
		EndTime:   time.Now(),
	}
	println("hello")
}
