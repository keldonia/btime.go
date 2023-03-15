# bTime

![Test Coverage](./coverage-badge.svg) [![Mutation testing badge](https://img.shields.io/endpoint?style=flat&url=https%3A%2F%2Fbadge-api.stryker-mutator.io%2Fgithub.com%2Fkeldonia%2FbTime%2Fmaster)](https://dashboard.stryker-mutator.io/reports/github.com/keldonia/bTime/master)

**bTime** is a small lightweight library, with few dependencies, primarily in test 
and 100% test and mutation test coverage, designed to help manage schedules using
bit manipulation.  It is particularly suited to working with dense schedules with 
discreet time intervals, e.g. 5 minutes.

## Getting Started

### From Source

1. Download the github repository

### Development and Testing

1. run `go test` to run all tests
1. run `go get -t -v github.com/avito-tech/go-mutesting/...` to pull go-mutesting, note you'll need to turn off go modules, as the module needs to be installed at go root.
1. run `go-mutesting ./core ./utils` to run mutation tests

### To Use in my Project
1. run `go get github.com/keldonia/btime.go`


## Using bTime

**bTime** includes two primary classes, [BScheduler](./documentation/scheduler.md) and [BTimeFactory](./documentation/bTimeFactory.md).  
Scheduler instantiates its own `BTimeFactory` when instantiated.  `BTimeFactory` 
can also be instantiated separately if one desires to directly make use of the btime utils;

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bScheduler, err := NewBScheduler(5); 

  // To get remaining availability
  remainingAvailabiltiy, err := bScheduler.GetCurrentAvailability(schedule);

  // To update a schedule
  updatedSchedule, err := bScheduler.UpdateSchedule(proposedSchedule, currentSchedule);

  // To process an appointment
  processedSchedule, err := bScheduler.ProcessAppointment(appointment, schedule, constants.BOOKING_UPDATE);

  // To process an array of appointments
  processedSchedule, err := bScheduler.ProcessAppointments(appointments, schedule, constants.BOOKING_UPDATE);

  // To convert a schedule to an appointment schedule
  processedSchedule, err := bScheduler.ConvertScheduleToAppointmentSchedule(schedule);

  // If using the factory directly
  // Instantiates a new BTimeFactory with a time interval of 5 min.
  bTimeFactory, err := NewBTimeFactory(5);  
```

Further information on bTime can be found below:
* [BScheduler](./documentation/scheduler.md)
* [BTimeFactory](./documentation/bTimeFactory.md)
* [How to use bTime](./documentation/howTos.md)

## Coming Features

1. Bit Array mode (don't store as strings), optional parameter
1. Allow compressed storage mode

## Assumptions

1. Weeks start on Sunday, this is the same as Golang time object and JS's Date object
1. Times are expected to be in UTC, utility converters are provided
1. Using 2016 * 2 digits (as a string, once for the schedule, once for bookings) to store the time for the week at 5 minute intervals, note this adjusts as time interval size changes.  The efficiency is greater for denser schedules.
