package leetcode

func calcScore(player []int) int {
	if len(player) == 0 {
		return 0
	}
	if len(player) == 1 {
		return player[0]
	}
	var score int
	if player[0] == 10 {
		score = player[0] + (2 * player[1])
	} else {

		score = player[0] + player[1]
	}

	n := len(player)

	for i := 2; i < n; i++ {
		if player[i-1] == 10 || player[i-2] == 10 {
			score += (2 * player[i])
		} else {
			score += player[i]
		}
	}
	return score
}

// question 2660
func IsWinner(player1, player2 []int) int {
	score1, score2 := calcScore(player1), calcScore(player2)

	if score1 > score2 {
		return 1
	}
	if score1 < score2 {
		return 2
	}
	return 0
}
