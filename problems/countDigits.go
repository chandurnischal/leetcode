package leetcode

// question 2520
func CountDigits(num int) int {
	nDup := num
	var count int
	for nDup != 0 {
		rem := nDup % 10
		nDup = nDup / 10
		if num%rem == 0 {
			count++
		}
	}
	return count
}
