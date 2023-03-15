# BTimeFactory

`BTimeFactory` is the other class that is available to end users.  It is exposed for users who desire to directly interact with the utilities provided.  As such `BTimeFactory` exposes several utility methods.

##  Instantiating

`BTimeFactory` only requires one argument in it's constructor, the time interval.  The time interval will affect the temporal resolution of the scheduler, ie if was set to 5 - the schedule's resolution is 5 minute intervals.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := NewBTimeFactory(5); 
```

**NB**: Each schedule group, ie all appointments for a given person, should have the same time interval

## `#ParseBString`

The `#ParseBString` method converts a bString in to a number so that it may be operated on, eg '10' = 2.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  convertedBString, err := bTimeFactory.ParseBString(bString);
```

## `#GenerateBString`

The `#GenerateBString` method converts an appointment into its bString representation. If the appointment is invalid, it return false.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  generatedBString, err := bTimeFactory.GenerateBString(appt);
```

## `#TimeStringSplit`

The `#TimeStringSplit` splits a bString into intervals dependent upon the time interval, each interval is one hour long.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  splitBStriing, err := bTimeFactory.TimeStringSplit(string);
```

## `#DecimalToBString`

The `#DecimalToBString` converts number into a bString representation with the given scheduling interval.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  splitBStriing, err := bTimeFactory.DecimalToBString(string);
```

## `#TestViabilityAndCompute`

The `#TestViabilityAndCompute` tests that two time intervals do not overlap, either returning the result of a bitwise OR function performed on the two numbers, or false if value returned from bitwise OR isn't equal to bitwise XOR.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  computedValue, err := bTimeFactory.TestViabilityAndCompute(bString1, bString2);
```

## `#DeleteAppointment`

The `#DeleteAppointment` tests removal a give time slot from a given time interval and if valid removes it from the scheduleInterval, else it returns the original scheduleInterval.

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  computedBString, err := bTimeFactory.DeleteAppointment(appointmentToDelete, scheduleInterval);
```

## `#ModifyScheduleAndBooking`

`#ModifyScheduleAndBooking` tests that an timeSlot does not overlap with another
timeSlot, if it does not overlap, the timeSlot is added to the bookings, else
return false.  Additionally, this method checks that the timeslot is within
availabilities (test).

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  computedBString, err := bTimeFactory.ModifyScheduleAndBooking(
    scheduleBStringToModify,
    scheduleBStringToTest,
    appt
  );
```

## `#ConvertScheduleToAppointmentSchedule`

`#ConvertScheduleToAppointmentSchedule` takes a schedule and availabilty converting
them into an array of appointments for each date

```golang
  // If using the scheduler
  // Instantiates a new Scheduler with a time interval of 5 min.
  bTimeFactory, err := new BTimeFactory(5); 

  appointmentSchedule := bTimeFactory.convertScheduleToAppointmentSchedule(
    schedule,
    availability
  );
```

## Additional Information

Additional information for each method and class is available in the form of JSDocs.