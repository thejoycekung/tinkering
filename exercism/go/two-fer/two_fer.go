// Twofer gives the message "One for X, one for me."
package twofer

// ShareWith takes an optional string `name` and returns the string "One for name, one for me."
// If name is not given, "you" is used.
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
