// Package raindrops allows a user to get a string of raindrops.
package raindrops

import "strconv"

// Convert takes an integer and returns a string indicating its raindrops.
// If there are no raindrops, it returns a string with the given integer.
func Convert(number int) string {
	var result string
	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}
