package main

// question 2078
func MaxDistance(colours []int) int {
	dist := 0
	n := len(colours)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if colours[i] != colours[j] {
				diff := j - i
				if diff > dist {
					dist = diff
				}
			}

		}
	}
	return dist
}
