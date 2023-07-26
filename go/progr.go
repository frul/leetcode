package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	sort.Ints(arr)
	d := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		d2 := arr[i] - arr[i-1]
		if d2 != d {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
