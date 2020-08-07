// Package hamming calculates Hamming distance.
package hamming

import "errors"

// Distance takes two same-length strings and returns their Hamming distance as an int.
// If the strings are not the same length, it returns an error.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("The given strings should be the same length")
	}

	var distance int = 0
	var bRune = []rune(b)
	for i, r := range []rune(a) {
		if r != bRune[i] {
			distance++
		}
	}

	return distance, nil
}
