package leetcode

import (
	"regexp"
	"strconv"
)

// question 2042
func AreNumbersAscending(str string) bool {
	re := regexp.MustCompile(`\d+`)
	match := re.FindAllString(str, -1)

	n := len(match)

	for i := 0; i < n-1; i++ {
		a, _ := strconv.Atoi(match[i])
		b, _ := strconv.Atoi(match[i+1])
		if a >= b {
			return false
		}
	}
	return true
}
