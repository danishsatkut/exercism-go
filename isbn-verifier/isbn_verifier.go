package isbn

import (
	"errors"
	"fmt"
	"strconv"
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

// ConvertToISBN13 converts an ISBN-10 to ISBN-13
func ConvertToISBN13(isbn10 string) (string, error) {
	if !IsValidISBN(isbn10) {
		return "", errors.New("invalid ISBN-10")
	}

	isbn10 = strings.ReplaceAll(isbn10, "-", "")
	isbn13 := "978" + isbn10

	return formatISBN13(isbn13, generateCheckDigitForISBN13(isbn13)), nil
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

func generateCheckDigitForISBN13(isbn string) string {
	checksum := 0

	for i, r := range isbn[0:12] {
		digit := int(r - '0')

		if i%2 == 0 {
			checksum += digit
		} else {
			checksum += digit * 3
		}
	}

	checkDigit := 10 - checksum%10
	if checkDigit == 10 {
		checkDigit = 0
	}

	return strconv.Itoa(checkDigit)
}

func formatISBN13(isbn, checkDigit string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", isbn[0:3], isbn[3:4], isbn[4:7], isbn[7:12], checkDigit)
}
