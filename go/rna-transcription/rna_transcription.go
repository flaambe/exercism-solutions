// Package strand provides an RNA transcription method
package strand

// ToRNA returns the RNA complement for a given DNA strand
func ToRNA(dna string) (result string) {
	for _, v := range dna {
		if v == 'G' {
			result += "C"
		} else if v == 'C' {
			result += "G"
		} else if v == 'T' {
			result += "A"
		} else if v == 'A' {
			result += "U"
		}
	}

	return
}
