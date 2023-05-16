package leetcode

func isBulky(length, width, height, mass int) bool {
	return length >= 1e4 || height >= 1e4 || width >= 1e4 || mass >= 1e4 || (length*width*height >= 1e9)
}

func isHeavy(mass int) bool {
	return mass >= 100
}

// question 2525
func CategorizeBox(length, width, height, mass int) string {
	bulk := isBulky(length, width, height, mass)
	heavy := isHeavy(mass)

	if bulk && heavy {
		return "Both"
	}
	if !bulk && !heavy {
		return "Neither"
	}
	if bulk && !heavy {
		return "Bulky"
	}
	return "Heavy"
}
