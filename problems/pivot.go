package leetcode

func leftSum(x int) int {
	var sum int
	for i := 1; i <= x; i++ {
		sum += i
	}
	return sum
}

func rightSum(x, n int) int {
	var sum int

	for i := x; i <= n; i++ {
		sum += i
	}

	return sum
}

func isPivot(x, n int) bool {
	return leftSum(x) == rightSum(x, n)
}

// question 2485
func PivotInteger(n int) int {
	for i := 1; i <= n; i++ {
		if isPivot(i, n) {
			return i
		}

	}
	return -1
}
