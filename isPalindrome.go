package main

import "strconv"

// question 9
func IsPalindrome(ele int) bool {
	str := strconv.Itoa(ele)
	n := len(str)
	var revStr string

	for i := n - 1; i >= 0; i-- {
		revStr += string(str[i])
	}
	return str == revStr
}
