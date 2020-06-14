package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	mu := sync.RWMutex{}
	m := FreqMap{}

	for _, text := range texts {
		go func(t string) {
			for _, r := range t {
				mu.Lock()
				m[r]++
				mu.Unlock()
			}
		}(text)
	}

	return m
}
