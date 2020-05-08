package romannumerals

import "errors"

// ToRomanNumeral converts from arabic numbers to Roman Numerals
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("out of range: " + string(n))
	}

	return convert(n), nil
}

var (
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = []string{"", "M", "MM", "MMM"}
)

func convert(n int) string {
	return thousands[n%1e4/1e3] + hundreds[n%1e3/1e2] + tens[n%1e2/1e1] + ones[n%1e1]
}
