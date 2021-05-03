// Package letter provides methods for counting the frequency of letters in texts.
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

func ConcurrentFrequency(s []string) FreqMap {
	res := make(FreqMap)
	ch := make(chan FreqMap, len(s))

	var wg sync.WaitGroup
	wg.Add(len(s))

	for _, v := range s {
		go func(s string) {
			ch <- Frequency(s)
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for m := range ch {
		for k, v := range m {
			res[k] += v
		}
	}

	return res
}
