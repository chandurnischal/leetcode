package leetcode

func signFunc(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

// question 1822
func ArraySign(nums []int) int {
	prod := 1
	for _, num := range nums {
		prod *= signFunc(num)
	}
	return prod
}
