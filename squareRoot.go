package main

// question 69
func Sqrt(n int) int {
	if n == 1 {
		return 1
	}
	low := 0
	high := n / 2

	for low <= high {
		mid := (low + high) / 2
		if mid*mid == n {
			return mid
		} else if mid*mid < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}
