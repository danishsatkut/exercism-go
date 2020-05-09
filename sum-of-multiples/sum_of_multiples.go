package summultiples

func SumMultiples(upto int, numbers ...int) int {
	uniqueMultiples := map[int]interface{}{}

	for _, number := range numbers {
		for _, m := range multiples(number, upto) {
			uniqueMultiples[m] = nil
		}
	}

	total := 0
	for m, _ := range uniqueMultiples {
		total += m
	}

	return total
}

func multiples(n int, upto int) []int {
	var result []int

	for i := 0; i < upto; i++ {
		if m := n * i; m < upto {
			result = append(result, m)
		}
	}

	return result
}
