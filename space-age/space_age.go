package space

type Planet string

const earthYearInSeconds = 60 * 60 * 24 * 365.25

func Age(s float64, p Planet) float64 {
	return s / earthYearInSeconds
}
