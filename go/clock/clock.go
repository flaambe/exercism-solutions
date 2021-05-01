// Package clock provides functionality for displaying time.
package clock

import "fmt"

// Clock represents time as minutes.
type Clock struct {
	minutes int
}

// New returns a clock that contains time in minutes.
func New(hour, minute int) Clock {
	mins := (hour*60 + minute) % 1440

	if mins < 0 {
		mins += 1440
	}

	return Clock{mins}
}

// String returns the time in hh:mm format.
func (c Clock) String() string {
	hours := c.minutes / 60
	minutes := c.minutes % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// Add adds the specified minutes to the current time.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract subtracts the specified minutes from the current time.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
