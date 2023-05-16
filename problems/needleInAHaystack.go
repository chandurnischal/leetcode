package leetcode

import "regexp"

// question 28
func NeedleInAHaystack(haystack, needle string) int {
	re := regexp.MustCompile(needle)
	match := re.FindSubmatchIndex([]byte(haystack))

	if len(match) == 0 {
		return -1
	}
	return match[0]
}
