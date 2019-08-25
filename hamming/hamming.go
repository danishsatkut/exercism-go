package hamming

import "fmt"

func Distance(strand, otherStrand string) (int, error) {
	var distance = 0

	if len(strand) != len(otherStrand) {
		return 0, fmt.Errorf("strands contain different number of nucleotides")
	}

	for i := 0; i < len(strand); i++ {
		if strand[i] != otherStrand[i] {
			distance += 1
		}
	}

	return distance, nil
}
