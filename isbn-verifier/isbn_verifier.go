package isbn

import (
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

	return isValidChecksum(isbn)
}

func isValidChecksum(isbn string) bool {
	checksum := 0

	for i, r := range isbn {
		if i == 9 && r == 'X' {
			checksum += 10
		} else {
			if !unicode.IsDigit(r) {
				return false
			}

			checksum += int(r - '0') * (10 - i)
		}
	}

	if checksum%11 == 0 {
		return true
	}

	return false
}
