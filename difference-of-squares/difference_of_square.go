package diffsquares

// Difference calculates the difference between the square of the sum and
// the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SumOfSquares calculates the sum of squares of first N natural numbers.
func SumOfSquares(n int) int {
	var total int

	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

// SquareOfSum calculates the square of sum of first N natural numbers.
func SquareOfSum(n int) int {
	var total int

	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}
