package leetcode

// question 2582
func PassThePillow(n, time int) int {
	quotient := time / (n - 1)
	remainder := time % (n - 1)
	if quotient%2 == 0 {
		return remainder + 1
	}
	return n - remainder
}
