package leetcode

import "sort"

// question 2570
func mergeArrays(nums1, nums2 [][]int) [][]int {
	sum := make(map[int]int)

	for _, v := range nums1 {
		sum[v[0]] += v[1]
	}

	for _, v := range nums2 {
		sum[v[0]] += v[1]
	}
	keys := []int{}

	for k := range sum {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var res [][]int

	for _, k := range keys {
		res = append(res, []int{k, sum[k]})
	}
	return res
}
