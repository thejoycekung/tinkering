// Package scrabble offers Scrabble-related tools.
package scrabble

import "unicode"

// Score calculates a given string's Scrabble score as an integer and returns it.
func Score(word string) int {
	var result int
	for _, r := range word {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}
	}
	return result
}
