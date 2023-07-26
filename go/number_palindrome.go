package main

import (
	"fmt"
	"strconv"
)

func reverseString(s string) (result string) {
	runes1 := []rune(s)
	runes := make([]rune, len(runes1))
	copy(runes, runes1)
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	result = string(runes)
	return
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	s2 := reverseString(s)
	fmt.Println(s)
	fmt.Println(s2)
	return s == s2
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-101))
}
