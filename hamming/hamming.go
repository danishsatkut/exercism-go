// Package hamming provides functionality to calculate hamming distance
package hamming

import (
	"fmt"
)

// Distance function calculates the hamming distance for two dna strands.
// It returns an error if the number of nucleotides differ between the two strands.
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
