package summultiples

func SumMultiples(upto int, numbers ...int) int {
	uniqueMultiples := map[int]interface{}{}

	for i := 0; i < upto; i++ {
		for _, number := range numbers {
			if multiple := number * i; multiple < upto {
				uniqueMultiples[multiple] = nil
			}
		}
	}

	total := 0
	for m, _ := range uniqueMultiples {
		total += m
	}

	return total
}
