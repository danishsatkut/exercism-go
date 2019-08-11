package space

import "fmt"

type Planet string

const earthYearInSeconds = 60 * 60 * 24 * 365.25

func Age(s float64, p Planet) float64 {
	switch p {
	case "Earth":
		return s / earthYearInSeconds
	case "Mercury":
		return s / earthYearInSeconds / 0.2408467
	case "Venus":
		return s / earthYearInSeconds / 0.61519726
	case "Mars":
		return s / earthYearInSeconds / 1.8808158
	case "Jupiter":
		return s / earthYearInSeconds / 11.862615
	case "Saturn":
		return s / earthYearInSeconds / 29.447498
	case "Uranus":
		return s / earthYearInSeconds / 84.016846
	case "Neptune":
		return s / earthYearInSeconds / 164.79132
	default:
		panic(fmt.Sprint("Invalid planet: ", p))
	}
}
