package summultiples

// SumMultiples calculates sum of all the unique multiples of particular
// numbers up to but not including that number.
func SumMultiples(upto int, numbers ...int) int {
	uniqueMultiples := map[int]struct{}{}

	for i := 0; i < upto; i++ {
		for _, number := range numbers {
			if multiple := number * i; multiple < upto {
				uniqueMultiples[multiple] = struct{}{}
			}
		}
	}

	total := 0
	for m := range uniqueMultiples {
		total += m
	}

	return total
}
