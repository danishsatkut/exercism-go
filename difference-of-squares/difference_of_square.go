package diffsquares

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func SumOfSquares(n int) int {
	var total int

	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

func SquareOfSum(n int) int {
	var total int

	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}

