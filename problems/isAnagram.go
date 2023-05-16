package leetcode

// question 242
func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	freq1, freq2 := make(map[rune]int), make(map[rune]int)
	chars := []rune{}

	for _, s := range str1 {
		_, ok := freq1[s]
		if !ok {
			chars = append(chars, s)
		}
		freq1[s]++
	}
	for _, s := range str2 {
		freq2[s]++
	}

	for _, char := range chars {
		if freq1[char] != freq2[char] {
			return false
		}
	}

	return true
}
