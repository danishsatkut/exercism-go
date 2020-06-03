package prime

import "math"

// Nth returns the prime number at nth position.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	for p := 2; ; p++ {
		if isPrime(p) {
			n--
			if n == 0 {
				return p, true
			}
		}
	}
}

func isPrime(p int) bool {
	if p%2 == 0 {
		return p == 2
	}

	max := int(math.Sqrt(float64(p))) + 1
	for i := 3; i < max; i += 2 {
		if p%i == 0 {
			return false
		}
	}

	return true
}
