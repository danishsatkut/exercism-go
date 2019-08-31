package space

import "fmt"

// Planet represents a planet
type Planet string

const earthYearInSeconds = 60 * 60 * 24 * 365.25

var planetEarthYears = map[Planet]float64{
	"Mercury": 0.2408467 * earthYearInSeconds,
	"Venus":   0.61519726 * earthYearInSeconds,
	"Earth":   1 * earthYearInSeconds,
	"Mars":    1.8808158 * earthYearInSeconds,
	"Jupiter": 11.862615 * earthYearInSeconds,
	"Saturn":  29.447498 * earthYearInSeconds,
	"Uranus":  84.016846 * earthYearInSeconds,
	"Neptune": 164.79132 * earthYearInSeconds,
}

// Age returns the age on a planet in earth years.
func Age(s float64, p Planet) float64 {
	if yearsInSecond, ok := planetEarthYears[p]; ok {
		return s / yearsInSecond
	}

	panic(fmt.Sprint("Invalid planet: ", p))
}
