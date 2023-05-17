package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func strToNum(str string) int {
	var res string
	alphabet := "abcdefghij"

	for _, s := range str {
		res += fmt.Sprintf("%d", strings.IndexRune(alphabet, s))
	}
	i, _ := strconv.Atoi(res)
	return i
}

// question 1880
func IsSumEqual(firstWord, secondWord, targetWord string) bool {
	return (strToNum(firstWord) + strToNum(secondWord)) == strToNum(targetWord)
}
