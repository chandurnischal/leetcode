package leetcode

// question 1720
func DecodeXOR(encoded []int, first int) []int {
	decoded := []int{first}

	for i, enc := range encoded {
		decoded = append(decoded, enc^decoded[i])
	}
	return decoded
}
