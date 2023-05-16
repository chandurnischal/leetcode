package leetcode

import (
	"fmt"
	"strconv"
)

// question 231
func IsPowerOf2(n int) bool {
	if n < 0 {
		return false
	}
	nBin := fmt.Sprintf("%b", n)
	sum := 0

	for _, b := range nBin {
		intB, _ := strconv.Atoi(string(b))
		sum += intB
	}
	return sum == 1
}
