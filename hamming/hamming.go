// Package hamming provides functionality to calculate hamming distance
package hamming

import (
	"errors"
)

// Distance function calculates the hamming distance for two dna strands.
// It returns an error if the number of nucleotides differ between the two strands.
func Distance(strand, otherStrand string) (int, error) {
	r1 := []rune(strand)
	r2 := []rune(otherStrand)

	if len(r1) != len(r2) {
		return 0, errors.New("strands contain different number of nucleotides")
	}

	var distance = 0

	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			distance++
		}
	}

	return distance, nil
}
