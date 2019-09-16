package strand

import "errors"

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
