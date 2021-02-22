package bob

import (
	"fmt"
	"regexp"
)

func main() {
	_ = regexp.MustCompile("^[^a-z]*\\!*\\s*$")
	_ = regexp.MustCompile("^[^a-zA-Z]*$")
	question := regexp.MustCompile("^.*\\?\\s*$")

	if question.MatchString("1, 2, 3?") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
