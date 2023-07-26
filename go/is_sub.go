package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	t_index := 0
	s_index := 0
	for t_index < len(t) && s_index < len(s) {
		dummy_string := string(s[s_index])
		new_t_indx := strings.Index(t[t_index:], dummy_string)
		if new_t_indx >= 0 {
			s_index++
			t_index = t_index + new_t_indx + 1
		} else {
			return false
		}
	}
	if s_index < len(s) {
		return false
	}
	return true
}

func main() {
	str1 := "acb"
	str2 := "ahbgdc"
	fmt.Println(isSubsequence(str1, str2))
}
