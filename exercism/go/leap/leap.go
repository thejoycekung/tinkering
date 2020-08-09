// Package leap offers leap year related utilities.
package leap

// IsLeapYear takes an int year and determines whether it is a leap year.
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	}

	return false
}
