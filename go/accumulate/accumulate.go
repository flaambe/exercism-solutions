package accumulate

// Accumulate returns a new collection containing the result of applying
// given operation to each element of the input collection.
func Accumulate(c []string, converter func(string) string) []string {
	result := make([]string, len(c))

	for i, v := range c {
		result[i] = converter(v)
	}

	return result
}
