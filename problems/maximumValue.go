package leetcode

import "strconv"

func valueOfString(str string) int {
	sum := ""

	for _, s := range str {
		_, err := strconv.Atoi(string(s))
		if err != nil {
			return len(str)
		}
		sum += string(s)
	}
	res, _ := strconv.Atoi(sum)
	return res
}

// question 2496
func MaximumValue(strs []string) int {
	maxVal := valueOfString(strs[0])

	for _, str := range strs {
		val := valueOfString(str)

		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
