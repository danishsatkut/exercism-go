package protein

import (
	"errors"
)

var (
	// ErrStop represents an error for STOP codon
	ErrStop = errors.New("terminating codon")

	// ErrInvalidBase represents an error for invalid RNA
	ErrInvalidBase = errors.New("invalid base")
)

// FromRNA converts RNA sequence into proteins
func FromRNA(rna string) ([]string, error) {
	var polypeptide = make([]string, 0, len(rna)/3)

	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				return polypeptide, nil
			}

			return polypeptide, err
		}

		polypeptide = append(polypeptide, protein)
	}

	return polypeptide, nil
}

// FromCodon converts codon into protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}
