package leetcode

import (
	"strings"
	"unicode"
)

func isLongEnough(str string) bool {
	return len(str) > 8
}

func hasLowercase(str string) bool {
	for _, s := range str {
		if unicode.IsLower(s) {
			return true
		}
	}
	return false
}

func hasUppercase(str string) bool {
	for _, s := range str {
		if unicode.IsUpper(s) {
			return true
		}
	}
	return false
}

func hasDigit(str string) bool {
	for _, s := range str {
		if unicode.IsDigit(s) {
			return true
		}
	}
	return false
}

func hasSpecial(str string) bool {

	special := "!@#$%^&*()-+"
	for _, s := range str {
		if strings.ContainsRune(special, s) {
			return true
		}
	}
	return false
}

func hasAdjacent(str string) bool {
	n := len(str)

	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

// question 2299
func StrongPasswordCheckerII(password string) bool {
	return isLongEnough(password) && hasLowercase(password) && hasUppercase(password) && hasDigit(password) && hasSpecial(password) && (!hasAdjacent(password))
}
