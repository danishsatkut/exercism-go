package clock

import (
	"fmt"
)

// New creates a clock with the specified hour and minutes
func New(hour, minutes int) Clock {
	hour = hour % 24
	if hour < 0 {
		hour = 24 + hour
	}

	c := Clock{hour, 0}
	if minutes < 0 {
		return c.Subtract(minutes * -1)
	}

	return c.Add(minutes)
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
	h := (minutes / 60) % 24
	m := minutes % 60

	m = c.minutes + m
	if m >= 60 {
		h++
		m = m - 60
	}

	h = c.hour + h
	if h >= 24 {
		h = h - 24
	}

	return Clock{h, m}
}

// Subtract subtracts specified minutes from the clock
func (c Clock) Subtract(minutes int) Clock {
	h := (minutes / 60) % 24
	m := minutes % 60

	m = c.minutes - m
	if m < 0 {
		h++
		m = 60 + m
	}

	h = c.hour - h
	if h < 0 {
		h = 24 + h
	}

	return Clock{h, m}
}
