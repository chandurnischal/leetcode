package leetcode

func matchesPlayerPerRound(teams int) (int, int) {
	n := teams / 2
	if teams%2 == 0 {
		return n, n
	}
	return n, n + 1
}

// question 1688
func NumberOfMatches(n int) int {
	var totalMatch int
	for n != 1 {
		match, teamsAdv := matchesPlayerPerRound(n)
		n = teamsAdv
		totalMatch += match
	}
	return totalMatch
}
