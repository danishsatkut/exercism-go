package clock

import (
	"strconv"
)

func New(hour, minutes int) Clock {
	if hour >= 24 {
		hour = hour % 24
	}

	if hour < 0 {
		hour = 24 + (hour % 24)
	}

	c := Clock{hour, 0}

	if minutes < 0 {
		return c.Subtract(minutes * -1)
	}

	return c.Add(minutes)
}

type Clock struct {
	hour, minutes int
}

func (c Clock) String() string {
	h := formattedTime(c.hour)
	m := formattedTime(c.minutes)

	return h + ":" + m
}

func (c Clock) Add(minutes int) Clock {
	h := minutes / 60
	m := minutes % 60

	m = c.minutes + m
	if m >= 60 {
		h += 1
		m = m - 60
	}

	h = c.hour + (h % 24)
	if h >= 24 {
		h = h - 24
	}

	return Clock{h, m}
}

func (c Clock) Subtract(minutes int) Clock {
	h := (minutes / 60) % 24
	m := minutes % 60

	m = c.minutes - m
	if m < 0 {
		h = h + 1

		m = 60 + m
	}

	h = c.hour - h
	if h < 0 {
		h = 24 + h
	}

	return Clock{h, m}
}

func formattedTime(n int) string {
	t := strconv.Itoa(n)
	if n < 10 {
		t = "0" + t
	}

	return t
}
