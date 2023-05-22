package leetcode

// question 383
func CanConstruct(ransom, magazine string) bool {
	magMap := make(map[rune]int)

	for _, m := range magazine {
		magMap[m]++
	}

	for _, r := range ransom {
		if _, ok := magMap[r]; ok {
			magMap[r]--
			if magMap[r] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
