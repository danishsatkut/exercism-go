// Package leap provides functionality to calculate leap years
package leap

// IsLeapYear returns true if the year is leap or false otherwise.
func IsLeapYear(year int) bool {
	return divisibleBy(year, 4) &&
		(!divisibleBy(year, 100) || divisibleBy(year, 400))
}

func divisibleBy(year, n int) bool {
	return year%n == 0
}
