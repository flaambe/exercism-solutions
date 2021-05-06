// Package proverb generates a proverbs
package proverb

// Proverb displays the full text of this proverbial rhyme.
func Proverb(rhyme []string) []string {
	result := make([]string, 0, len(rhyme))

	for i := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, "And all for the want of a "+rhyme[0]+".")
			break
		}

		result = append(result, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}

	return result
}
