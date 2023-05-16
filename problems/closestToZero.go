package leetcode

func distToZero(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// question 2239
func FindClosestNumber(nums []int) int {
	minVal := nums[0]
	minDist := distToZero(minVal)

	for _, n := range nums {
		dist := distToZero(n)
		if dist < minDist {
			minDist = dist
			minVal = n
		}
		if dist == minDist {
			minVal = max(minVal, n)
		}
	}
	return minVal
}
