package leetcode

import (
	"strings"
)

func PrefixCount(words []string, pref string) int {
	var count int

	for _, word := range words {
		if strings.Index(word, pref) == 0 {
			count++
		}
	}
	return count
}
