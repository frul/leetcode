package main

import "fmt"

func largestAltitude(gain []int) int {
	result := 0
	sum := 0
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		if sum > result {
			result = sum
		}
	}
	return result
}

func main() {
	nums := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(nums))
}
