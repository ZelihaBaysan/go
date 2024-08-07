package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	Date, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return Date.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	Date, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return Date.Hour() >= 12 && Date.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	Date, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("You have an appointment on %s", Date.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedDate, err := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", time.Now().Year()))
	if err != nil {
		panic(err)
	}
	return parsedDate
}
