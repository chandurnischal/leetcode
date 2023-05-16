package leetcode

var memo = make(map[int]int)

// question 70
func ClimbStairs(n int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	memo[n] = ClimbStairs(n-1) + ClimbStairs(n-2)
	return memo[n]
}
