// Package isogram provides a method for determining whether a word or phrase is an isogram
package isogram

import "strings"

// IsIsogram determine if a word or phrase is an isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	visits := make(map[rune]bool, len(word))

	for _, v := range word {
		if v == '-' || v == ' ' {
			continue
		}

		if _, ok := visits[v]; ok {
			return false
		}

		visits[v] = true
	}

	return true
}
