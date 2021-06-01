// Package leap provides a method for determining whether a year is a leap year
package leap

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}

	return year%4 == 0
}
