package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(sentence string) Frequency {
	f := Frequency{}

	for _, word := range strings.Split(sentence, " ") {
		f[word]++
	}

	return f
}
