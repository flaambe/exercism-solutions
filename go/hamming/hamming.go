// Package hamming provides a method for calculating the Hamming Distance between two DNA strands
package hamming

import "errors"

// Distance returns the number of differences between two DNA strands
// If two sequences are of different length, method returns an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the two sequences have different lengths")
	}

	var result int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
