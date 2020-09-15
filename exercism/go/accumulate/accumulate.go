package accumulate

import "fmt"

// Accumulate takes a list of words and a function and applies the function to the words, returning the modified list.
func Accumulate(words []string, op func(string) string) []string {
	for _, str := range words {
		str = op(str)
		fmt.Println(str)
	}
	return words
}
