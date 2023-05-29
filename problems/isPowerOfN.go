package leetcode

// question 326
func isPowerOfN(num, n int) bool {
	if num == 1 {
		return true
	}
	if num < 1 {
		return false
	}
	if num%n != 0 {
		return false
	}
	return isPowerOfN(num/n, n)
}
