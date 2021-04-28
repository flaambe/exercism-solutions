// Package luhn provides a method for validating various identification numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if the given string is valid per the Luhn formula.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	parity := len(s) % 2
	sum := 0

	if len(s) <= 1 {
		return false
	}

	for i, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
		digit := int(v - '0')

		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return (sum % 10) == 0
}
