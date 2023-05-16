package leetcode

import "fmt"

// question 693
func HasAlternatingBits(n int) bool {
	binN := fmt.Sprintf("%b", n)
	l := len(binN)

	for i := 0; i < l-1; i++ {
		if binN[i] == binN[i+1] {
			return false
		}
	}
	return true
}
