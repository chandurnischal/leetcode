package main

import (
	"fmt"
	"strings"
)

// func CanConstructRec(target string, wordBank []string) bool {
// 	if target == "" {
// 		return true
// 	}

// 	for _, word := range wordBank {
// 		index := strings.Index(target, word)
// 		if index == 0 {

// 		}
// 	}
// }

func main() {
	target := "abcdef"
	wordBank := []string{"ab", "abc", "cd", "def", "abcd"}
	index := strings.Index(target, wordBank[0])

	fmt.Println(target[index:])

}
