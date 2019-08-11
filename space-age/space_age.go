package space

import "fmt"

type Planet string

const earthYearInSeconds = 60 * 60 * 24 * 365.25

var planetEarthYears = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(s float64, p Planet) float64 {
	if earthYears, ok := planetEarthYears[p]; ok {
		return s / (earthYearInSeconds * earthYears)
	}

	panic(fmt.Sprint("Invalid planet: ", p))
}
