package leetcode

// question 258
func AddDigits(n int) int {
	if n < 10 {
		return n
	}
	if n%9 == 0 {
		return 9
	}
	return n % 9
}
