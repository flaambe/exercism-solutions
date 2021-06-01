// Package strand provides an RNA transcription method
package strand

// ToRNA returns the RNA complement for a given DNA strand.
func ToRNA(dna string) (rna string) {
	for _, v := range dna {
		switch v {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}

	return rna
}
