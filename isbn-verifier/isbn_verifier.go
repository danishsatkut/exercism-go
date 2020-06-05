package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks the validity of the specified ISBN.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	return isValidChecksum(isbn)
}

func IsValidISBN13(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 13 {
		return false
	}

	return isValidISBN13Checksum(isbn)
}

func isValidChecksum(isbn string) bool {
	checksum := 0

	for i, r := range isbn {
		if unicode.IsDigit(r) {
			checksum += int(r-'0') * (10 - i)
		} else if i == 9 && r == 'X' {
			checksum += 10
		} else {
			return false
		}
	}

	return checksum%11 == 0
}

func isValidISBN13Checksum(isbn string) bool {
	checksum, err := calculateISBN13Checksum(isbn)
	if err != nil {
		return false
	}

	return checksum%10 == 0
}
