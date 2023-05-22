package leetcode

// question 2413
func SmallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
