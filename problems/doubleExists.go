package leetcode

// question 1346
func CheckIfExists(nums []int) bool {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && nums[j] == 2*nums[i] {
				return true
			}
		}
	}
	return false
}
