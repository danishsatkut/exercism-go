// Package strand provides functionality related to DNA and RNA strands.
package strand

import "strings"

var rnaTranscriptions = strings.NewReplacer(
	"G", "C",
	"C", "G",
	"T", "A",
	"A", "U",
)

// ToRNA converts a DNA strand into a RNA strand
func ToRNA(dna string) string {
	return rnaTranscriptions.Replace(dna)
}
