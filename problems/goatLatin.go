package leetcode

import "strings"

func startsWithVowel(str string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, v := range vowels {
		if v == rune(str[0]) {
			return true
		}
	}
	return false
}

func removeFirstCharAndAppendToLast(str string) string {
	temp := str[0]
	res := str[1:] + string(temp)
	return res
}

func appendMA(str string) string {
	return str + "ma"
}

func appendAs(str string, index int) string {
	temp := strings.Repeat("a", index+1)
	return str + temp
}

// question 824
func ToGoatLatin(sentence string) string {
	strSlice := strings.Split(sentence, " ")
	n := len(strSlice)
	res := ""
	for i, s := range strSlice {

		if !startsWithVowel(s) {
			s = removeFirstCharAndAppendToLast(s)
		}
		s = appendMA(s)
		s = appendAs(s, i)
		res += s
		if i != n-1 {
			res += " "
		}
	}
	return res
}
