package clock

import (
	"fmt"
)

// Clock represents a clock with hour and minutes
type Clock struct {
	minutes int
}

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

	return Clock{hour*60 + minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add adds specified minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract subtracts specified minutes from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
