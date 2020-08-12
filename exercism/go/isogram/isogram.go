// Package isogram offers utilities to determine if strings are isograms.
package isogram

import "unicode"

// IsIsogram takes a word and determines whether it is an isogram, meaning it has no repeating letters.
func IsIsogram(word string) bool {
	var exists = make(map[rune]bool)
	for _, r := range word {
		if r == ' ' || r == '-' {
			continue
		}
		r = unicode.ToUpper(r)
		if exists[r] {
			return false
		}
		exists[r] = true
	}
	return true
}
