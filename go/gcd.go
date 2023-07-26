package main

import "fmt"

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func isStringNDivisible(str string, n int) bool {
	len := len(str)
	if len%n != 0 {
		return false
	}
	i := 0
	for i < (len/n)-1 {
		for j := i; j < i+n; j++ {
			if str[j] != str[j+n] {
				return false
			}
		}
		i += n
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)
	n := gcd(n1, n2)

	for i := n; i >= 1; i-- {
		if isStringNDivisible(str1, i) {
			if isStringNDivisible(str2, i) {
				if str1[:i] == str2[:i] {
					return str2[:i]
				}
			}
		}
	}

	return ""
}

func main() {
	str1 := "LEFT"
	str2 := "CODE"
	fmt.Println(gcdOfStrings(str1, str2))
}
