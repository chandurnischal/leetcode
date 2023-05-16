package leetcode

// question 169
func MajorityElement(nums []int) int {
	n := len(nums)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for k, v := range freq {
		if v > (n / 2) {
			return k
		}
	}
	return -1
}
