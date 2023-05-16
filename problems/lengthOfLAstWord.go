package leetcode

import "regexp"

// question 58
func LengthOfLastWord(str string) int {
	pattern := regexp.MustCompile("[a-zA-Z]+")

	match := pattern.FindAllStringIndex(str, -1)
	n := len(match) - 1
	return match[n][1] - match[n][0]
}
