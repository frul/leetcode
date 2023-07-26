package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max_avg := float64(sum) / float64(k)
	for i := 1; i+k <= len(nums); i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		avg := float64(sum) / float64(k)
		if avg > max_avg {
			max_avg = avg
		}
	}
	return max_avg
}

func main() {
	nums := []int{0, 1, 1, 3, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
