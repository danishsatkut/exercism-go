package encode

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {
	var encoded string

	var previous rune
	var count = 1

	for i, current := range input {
		if current == previous {
			count++
		} else
		{
			if i != 0 {
				if count > 1 {
					encoded += strconv.Itoa(count)
				}

				encoded += fmt.Sprintf("%c", previous)
				count = 1
			}
		}

		if i == len(input) - 1 {
			if count > 1 {
				encoded += strconv.Itoa(count)
			}

			encoded += fmt.Sprintf("%c", current)
		}

		previous = current
	}

	return encoded
}
