// Package strand provides functionality related to DNA and RNA strands.
package strand

var rnaTranscriptions = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA strand into a RNA strand
func ToRNA(dna string) string {
	var rna string

	for _, nucleotide := range dna {
		n, ok := rnaTranscriptions[string(nucleotide)]
		if !ok {
			panic("invalid DNA dnaNucleotide: " + string(nucleotide))
		}

		rna += n
	}

	return rna
}
