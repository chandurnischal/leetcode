package leetcode

import (
	"regexp"
	"strings"
)

// question 434
func CountSegments(str string) int {
	// trim whitespaces in the edges
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return 0
	}

	// replace multiple spaces within text with one whitespace
	re := regexp.MustCompile(`\s{2,}`)
	str = re.ReplaceAllString(str, " ")

	// split the string using whitespaces
	strSlice := strings.Split(str, " ")

	return len(strSlice)
}
