package leetcode

import "strings"

// question 1108
func DefangIP(address string) string {
	address = strings.ReplaceAll(address, ".", "[.]")
	return address
}
