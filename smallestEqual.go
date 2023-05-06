package main

// question 2057
func SmallestEqual(nums []int) int {
	for i := range nums {
		if i%10 == nums[i] {
			return i
		}
	}
	return -1
}
