package leetcode

import "fmt"

// question 506
func FindRelativeRanks(score []int) []string {
	scoreN := []int{}
	indices := []int{}
	idx := make(map[int]int)

	for i, v := range score {
		scoreN = append(scoreN, v)
		indices = append(indices, i)
	}

	n := len(score)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if scoreN[i] < scoreN[j] {
				scoreN[i], scoreN[j] = scoreN[j], scoreN[i]
			}
		}
	}

	for i, v := range scoreN {
		idx[v] = indices[i]
	}
	var answer []string
	for _, v := range score {
		if idx[v] == 0 {
			answer = append(answer, "Gold Medal")
		} else if idx[v] == 1 {
			answer = append(answer, "Silver Medal")
		} else if idx[v] == 2 {
			answer = append(answer, "Bronze Medal")
		} else {
			answer = append(answer, fmt.Sprintf("%d", idx[v]+1))
		}
	}
	return answer
}
