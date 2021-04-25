// Package scrabble provides a method for calculating the Scrabble score
package scrabble

import "unicode"

// Score computes the Scrabble score for the given word
func Score(s string) int {
	var result int

	scrabble_scores := map[rune]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'l': 1,
		'n': 1,
		'r': 1,
		's': 1,
		't': 1,

		'd': 2,
		'g': 2,

		'b': 3,
		'c': 3,
		'm': 3,
		'p': 3,

		'f': 4,
		'h': 4,
		'v': 4,
		'w': 4,
		'y': 4,

		'k': 5,

		'j': 8,
		'x': 8,

		'q': 10,
		'z': 10,
	}

	for _, v := range s {
		result += scrabble_scores[unicode.ToLower(v)]
	}

	return result
}
