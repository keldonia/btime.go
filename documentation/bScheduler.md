# BScheduler

`BScheduler` represents the primary class that most users of **bTime** will use,
thus the primary interface users interact.  As such it is entirely self-contained
from the user perspective.  The primary methods of interaction with the class
are the methods: `#UpdateSchedule` and `#ProcessAppointment`;

##  Instantiating

`BScheduler` only requires one argument in it's constructor, the time interval.
The time interval will affect the temporal resolution of the scheduler, ie if it
was set to 5 - the schedule's resolution is 5 minute intervals.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bScheduler, err := NewBScheduler(5); 
```

**NB**: Each schedule group, ie all appointments for a given person,
should have the same time interval

## `#GetCurrentAvailability`

The `#GetCurrentAvailability` Takes a valid schedule and computes the remaining
availability based on the total availability and current bookings, throws an
error if an invalid scehdule is passed.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  scheduler, err := NewBScheduler(5); 

  // To get remaining availabiltiy in a schedule
  remainingAvailability, err := bScheduler.GetCurrentAvailability(schedule);
```

## `#UpdateSchedule`

The `#UpdateSchedule` method takes two arguments, the proposed schedule and the
current schedule.  These schedules must adhere to the Schedule interface.
It will compare the proposed schedule against the bookings of the current schedule
and will either return the updated schedule complete with the current bookings,
in the case that current bookings work with the updated schedule or throws an
error if they do not.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  scheduler, err := NewBScheduler(5); 

  // To update a schedule
  updatedSchedule, err := scheduler.UpdateSchedule(
    proposedSchedule,
    currentSchedule
  );
```

## `#ProcessAppointment`

The `#ProcessAppointment` method takes three agruments, the proposed appointment,
the schedule, and the type of action \- a booking update or appointment delete.
The method will then check if there is availability for the proposed appointment
in the schedule in the case of a booking update, returning with the updated schedule
if the appointment is compatible with the schedule, or throws an error if not.
If the an appointment is to be deleted the time interval of the appointment is freed.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bScheduler, err := NewBScheduler(5); 

  // To process an appointment
  processedSchedule, err := bScheduler.ProcessAppointment(
    appointment,
    schedule,
    ScheduleActions.BOOKING_UPDATE
  ); 
```

## `#ProcessAppointments`

The `#ProcessAppointments` method takes three agruments, the proposed appointments
\- in the form of an array, the schedule, and the type of action \- a booking update
or appointment delete.  The method will then check if there is availability for the
proposed appointments in the schedule in the case of a booking update, returning with
the updated schedule if the appointments are compatible with the schedule, or throws an 
error if not. If the appointments are to be deleted the time intervals of the appointments
are freed.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bScheduler, err := NewBScheduler(5); 

  // To process an appointment
  processedSchedule, err := bScheduler.ProcessAppointments(
    appointments, 
    schedule, 
    ScheduleActions.BOOKING_UPDATE
  ); 
```

## `#ConvertScheduleToAppointmentSchedule`

`#ConvertScheduleToAppointmentSchedule` takes a schedule and converts it into an
array of appointments for each date

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bScheduler, err := NewBScheduler(5); 

  appointmentSchedule, err := bScheduler.ConvertScheduleToAppointmentSchedule(
    schedule
  );
```

## Additional Information

Additional information for each method and class is available in the form of JSDocs.