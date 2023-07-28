package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mergeAlternately(word1 string, word2 string) string {
	runes := make([]rune, len(word1)+len(word2))
	i := 0
	j := 0
	for i < min(len(word1), len(word2)) {
		if j%2 == 0 {
			runes[j] = rune(word1[i])
		} else {
			runes[j] = rune(word2[i])
			i++
		}
		j++
	}
	for i < len(word1) {
		runes[j] = rune(word1[i])
		i++
		j++
	}
	for i < len(word2) {
		runes[j] = rune(word2[i])
		i++
		j++
	}
	return string(runes)
}

func main() {
	s1 := "abc"
	s2 := "pqr"
	fmt.Println(mergeAlternately(s1, s2))
}
