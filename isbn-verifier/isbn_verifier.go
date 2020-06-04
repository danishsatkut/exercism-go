package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	if unicode.IsLetter(rune(isbn[9])) && isbn[9] != 'X' {
		return false
	}

	checksum := 0
	for i, r := range isbn {
		if i == 9 && r == 'X' {
			checksum += 10
		} else {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return false
			}

			checksum += v * (10 - i)
		}
	}

	if checksum%11 == 0 {
		return true
	}

	return false
}
