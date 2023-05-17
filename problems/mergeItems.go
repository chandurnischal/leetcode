package leetcode

// question 2363
func MergeSimilarItems(items1, items2 [][]int) [][]int {
	weight := make(map[int]int)

	for _, item := range items1 {
		weight[item[0]] += item[1]
	}

	for _, item := range items2 {
		weight[item[0]] += item[1]
	}
	ret := [][]int{}

	for k, v := range weight {
		ret = append(ret, []int{k, v})
	}

	n := len(ret)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if ret[i][0] > ret[j][0] {
				ret[i], ret[j] = ret[j], ret[i]
			}
		}
	}
	return ret
}
