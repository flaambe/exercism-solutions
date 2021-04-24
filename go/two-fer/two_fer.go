// Package twofer provides a method for sharing something
package twofer

// ShareWith returns a string with the sharing message based on the specified name
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	}

	return "One for you, one for me."
}
