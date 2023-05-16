package leetcode

func strictlyLesser(nums []int, n int) bool {
	for _, num := range nums {
		if num < n {
			return true
		}
	}
	return false
}

func strictlyGreater(nums []int, n int) bool {
	for _, num := range nums {
		if num > n {
			return true
		}
	}
	return false
}

// question 2148
func CountElements(nums []int) int {
	var count int
	for _, num := range nums {
		if strictlyGreater(nums, num) && strictlyLesser(nums, num) {
			count++
		}
	}
	return count
}
