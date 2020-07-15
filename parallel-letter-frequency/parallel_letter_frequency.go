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

// ConcurrentFrequency concurrently counts the frequency of each rune in a given text
//  and returns this data as a FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	var (
		m  = FreqMap{}
		mu = sync.RWMutex{}
		wg = sync.WaitGroup{}
	)

	for _, text := range texts {
		wg.Add(1)

		go func(t string) {
			defer wg.Done()

			for _, r := range t {
				mu.Lock()
				m[r]++
				mu.Unlock()
			}
		}(text)
	}

	wg.Wait()

	return m
}
