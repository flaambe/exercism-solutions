// Package dna provides a method for counting nucleotide in DNA strand
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, v := range d {
		if strings.ContainsRune("ACGT", v) {
			h[v]++
		} else {
			return Histogram{}, errors.New("invalid DNA strand")
		}
	}

	return h, nil
}
