package leetcode

import "sort"

func inSlice(slice []int, ele int) bool {
	for _, s := range slice {
		if s == ele {
			return true
		}
	}
	return false
}

// question 2248
func Intersection(nums [][]int) []int {
	num := nums[0]
	count := make(map[int]int)
	intersect := []int{}
	n := len(nums)
	for _, ele := range num {
		for i := 0; i < n; i++ {
			if inSlice(nums[i], ele) {
				count[ele]++
			}
		}
	}

	for k, v := range count {
		if v == n {
			intersect = append(intersect, k)
		}
	}
	sort.Ints(intersect)
	return intersect
}
