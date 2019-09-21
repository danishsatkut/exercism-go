// Package dna provides functionalities related to DNA strands
package dna

import "errors"

// DNA is a list of nucleotides.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := newHistogram()

	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; !ok {
			return nil, errors.New("invalid nucleotide")
		}

		h[nucleotide]++
	}

	return h, nil
}

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

func newHistogram() Histogram {
	return Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
}
