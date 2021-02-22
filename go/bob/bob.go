// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// Hey returns Bob's limited response
func Hey(remark string) string {
	cap := regexp.MustCompile("^[^a-z]*\\!?\\s*$")
	letters := regexp.MustCompile("^[^a-zA-z]*[^/?!]?$")
	question := regexp.MustCompile("^.*\\?\\s*$")
	if b, _ := regexp.MatchString("^\\s*$", remark); b {
		return "Fine. Be that way!"
	} else if question.MatchString(remark) {
		if b, _ := regexp.MatchString("^[^a-z1-9:)]*$", remark); b {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if letters.MatchString(remark) {
		return "Whatever."
	} else if cap.MatchString(remark) {
		// If remark is a yell
		return "Whoa, chill out!"
	}
	return "Whatever."
}
