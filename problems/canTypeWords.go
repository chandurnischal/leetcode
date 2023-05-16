package leetcode

import "strings"

func canTypeWord(word, broken string) bool {
	for _, b := range broken {
		if strings.ContainsRune(word, b) {
			return false
		}
	}
	return true
}

// question 1935
func CanBeTypedWords(text, brokenLetters string) int {
	textSlice := strings.Split(text, " ")
	var count int

	for _, t := range textSlice {
		if canTypeWord(t, brokenLetters) {
			count++
		}
	}

	return count
}
