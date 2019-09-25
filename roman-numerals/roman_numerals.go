package romannumerals

import "errors"

func ToRomanNumeral(arabic int) (string, error) {
	if arabic == 0 {
		return "", errors.New("cannot convert 0 to roman numeral")
	}

	var roman string

	m := arabic % 5
	if m > 0 && m < 4 {
		for i := m; i > 0; i-- {
			roman += "I"
		}
	} else {
		if m != 0 {
			for i := m; i < 5; i++ {
				roman += "I"
			}
		}

		roman += "V"
	}

	return roman, nil
}
