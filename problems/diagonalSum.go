package leetcode

// question 1572
func DiagonalSum(mat [][]int) int {

	n := len(mat)
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || n-1-i == j {
				sum += mat[i][j]
			}
		}
	}
	return sum
}
