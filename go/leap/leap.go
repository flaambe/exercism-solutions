// Package leap provides a method for determining whether a year is a leap year
package leap

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 != 0 {
		return false
	}

	return true
}
