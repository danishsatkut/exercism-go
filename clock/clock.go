package clock

import (
	"fmt"
)

// Clock represents a clock using minutes
type Clock struct {
	minutes int
}

// New creates a clock with the specified hour and minutes
func New(hour, minutes int) Clock {
	m := minutes + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}

	return Clock{m}
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
