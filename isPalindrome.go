package main

import (
	"strconv"
	"strings"
	"unicode"
)

// question 9
func IsPalindromeInt(ele int) bool {
	str := strconv.Itoa(ele)
	n := len(str)
	var revStr string

	for i := n - 1; i >= 0; i-- {
		revStr += string(str[i])
	}
	return str == revStr
}

func removeSpecial(str string) string {
	var res string

	for _, s := range str {
		if unicode.IsLetter(s) || unicode.IsNumber(s) {
			res += strings.ToLower(string(s))
		}
	}

	return res
}

// question 125
func IsPalindromeStr(str string) bool {
	str = removeSpecial(str)
	n := len(str)
	var res string
	for i := n - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return str == res
}
