package leetcode

// question 14
func LongestPrefix(strs []string) string {
	res := ""
	n := len(strs[0])

	for i := 0; i < n; i++ {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res

}
