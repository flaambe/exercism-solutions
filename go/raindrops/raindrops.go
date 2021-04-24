// Package raindrops provides a method for converting a number to raindrop sounds
package raindrops

import "strconv"

// Convert convert a number into a string that contains raindrop sounds
// corresponding to certain potential factors
func Convert(a int) string {
	var result string

	if a%3 == 0 {
		result += "Pling"
	}

	if a%5 == 0 {
		result += "Plang"
	}

	if a%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(a)
	}

	return result
}
