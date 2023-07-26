package main

import "fmt"

func myPow(x float64, n int) float64 {
	result := 1.0
	if n >= 0 {
		for i := 0; i < n; i++ {
			result *= x
		}
	} else {
		for i := 0; i < -n; i++ {
			result *= x
		}
		result = 1 / result
	}
	return result
}

func main() {
	fmt.Println(myPow(2, -2))
}
