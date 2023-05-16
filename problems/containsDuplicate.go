package leetcode

// question 217
func ContainsDuplicates(nums []int) bool {
	count := make(map[int]int)

	for _, num := range nums {
		if _, ok := count[num]; ok {
			return true
		}
		count[num] = 1
	}

	return false
}
