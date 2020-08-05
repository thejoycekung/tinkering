// Bob is a snippy teenager who will talk to you.
package bob

import (
	"regexp"
	"strings"
)

// Hey will take a remark and give a response that is typical Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var is_alpha, _ = regexp.MatchString("[a-zA-Z]+", remark)
	var is_all_caps = strings.ToUpper(remark) == remark && is_alpha
	var is_question = remark != "" && strings.HasSuffix(remark, "?")

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case is_all_caps && is_question:
		return "Calm down, I know what I'm doing!"
	case is_all_caps:
		return "Whoa, chill out!"
	case is_question:
		return "Sure."
	default:
		return "Whatever."
	}
}
