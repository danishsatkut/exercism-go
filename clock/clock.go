package clock

import (
	"fmt"
)

// New creates a clock with the specified hour and minutes
func New(hour, minutes int) Clock {
	hour = hour + (minutes / 60)
	minutes = minutes % 60

	if minutes < 0 {
		hour--
		minutes = 60 + minutes
	}

	hour = hour % 24
	if hour < 0 {
		hour = 24 + hour
	}

	return Clock{hour, minutes}
}

// Clock represents a clock with hour and minutes
type Clock struct {
	hour, minutes int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minutes)
}

// Add adds specified minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minutes+minutes)
}

// Subtract subtracts specified minutes from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minutes-minutes)
}
