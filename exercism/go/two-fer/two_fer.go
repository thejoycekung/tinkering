// Package twofer gives the message "One for X, one for me." to a user.
package twofer

import "fmt"

// ShareWith takes an optional string `name` and returns the string "One for name, one for me." to the user.
// If `name` is not given, "you" is used.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
