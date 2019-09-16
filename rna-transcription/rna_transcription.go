// Package strand provides functionality related to DNA and RNA strands.
package strand

import "errors"

// ToRNA converts a DNA strand into a RNA strand
func ToRNA(dna string) string {
	var rna string

	for _, nucleotide := range dna {
		n, err := toRNANucleotide(string(nucleotide))
		if err != nil {
			panic(err)
		}

		rna += n
	}

	return rna
}

func toRNANucleotide(dnaNucleotide string) (string, error) {
	switch dnaNucleotide {
	case "C":
		return "G", nil
	case "G":
		return "C", nil
	case "T":
		return "A", nil
	case "A":
		return "U", nil
	default:
		return "", errors.New("invalid DNA nucleotide: " + dnaNucleotide)
	}
}
