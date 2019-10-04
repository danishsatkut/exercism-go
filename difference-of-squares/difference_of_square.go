package diffsquares

// Difference calculates the difference between the square of the sum and
// the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SumOfSquares calculates the sum of squares of first N natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum calculates the square of sum of first N natural numbers.
func SquareOfSum(n int) int {
	sum := sumOf(n)

	return sum * sum
}

func sumOf(n int) int {
	m := float64(n)

	return int((m / 2) * (m + 1))
}
