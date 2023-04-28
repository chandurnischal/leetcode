package main

import (
	"fmt"
	"regexp"
)

func findLongestLengthForChar(char rune, str string) int {
	strPat := fmt.Sprintf("%c(.+)?%c", char, char)
	pattern := regexp.MustCompile(strPat)
	match := pattern.FindStringSubmatch(str)
	if len(match) == 0 {
		return -1
	}
	if len(match[1]) == 0 {
		return 0
	}
	return len(match[1])
}

func MaxLengthBetweenEqualCharacters(str string) int {
	charCount := make(map[rune]int)

	for _, s := range str {
		charCount[s] = findLongestLengthForChar(s, str)
	}
	counts := []int{}

	for _, v := range charCount {
		counts = append(counts, v)
	}
	max := -1
	for _, c := range counts {
		if c > max {
			max = c
		}
	}

	return max
}
