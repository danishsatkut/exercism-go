package prime

// Nth returns the prime number at nth position.
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := make([]int, 0, n)
	primes = append(primes, 2)

	for i := 3; len(primes) <= n; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
	}

	return primes[n-1], true
}

func isPrime(primes []int, number int) bool {
	for _, prime := range primes {
		if number%prime == 0 {
			return false
		}
	}

	return true
}
