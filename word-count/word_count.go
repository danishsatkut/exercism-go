package wordcount

import (
	"regexp"
	"strings"
)

// Frequency represents count of occurrences of each word
type Frequency map[string]int

var (
	splitter      = regexp.MustCompile(`[\s\,\:]+`)
	wordExtractor = regexp.MustCompile(`([\w\'])+`)
	quotedWord    = regexp.MustCompile(`^\'(\w+)\'$`)
)

// WordCount calculates the occurrence of each word in a phrase
func WordCount(sentence string) Frequency {
	f := Frequency{}

	words := splitter.Split(sentence, -1)

	for _, word := range words {
		if w := wordExtractor.FindString(word); w != "" {
			if quotedWord.MatchString(w) {
				w = quotedWord.FindStringSubmatch(w)[1]
			}

			f[strings.ToLower(w)]++
		}
	}

	return f
}
