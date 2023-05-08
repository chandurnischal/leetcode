package main

// question 1
func TwoSum(nums []int, target int) []int {

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// question 1
func TwoSumV2(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if _, ok := prevMap[diff]; !ok {
			prevMap[num] = i
		} else {
			return []int{prevMap[diff], i}
		}
	}
	return nil
}
