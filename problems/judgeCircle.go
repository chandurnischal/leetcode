package leetcode

// question 657
func JudgeCircle(moves string) bool {
	destination := []int{0, 0}

	for _, move := range moves {
		if move == 'U' {
			destination[1]++
		}
		if move == 'D' {
			destination[1]--
		}
		if move == 'R' {
			destination[0]++
		}
		if move == 'L' {
			destination[0]--
		}
	}
	return destination[0] == 0 && destination[1] == 0
}
