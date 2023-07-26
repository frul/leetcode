package main

import (
	"fmt"
	"math"
)

func getAverages(nums []int, k int) []int {
	size := len(nums)
	k_sum := 0
	result := make([]int, size)
	for i := 0; i < size; i++ {
		if i-k < 0 || i+k >= size {
			result[i] = -1
		} else {
			if i-k == 0 {
				for j := 0; j <= i+k; j++ {
					k_sum += nums[j]
				}
			} else {
				k_sum -= nums[i-k-1]
				k_sum += nums[i+k]
			}
			avg := float64(k_sum) / float64(2*k+1)
			result[i] = int(math.Floor(avg))
		}
	}
	return result
}

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	fmt.Println(getAverages(nums, k))
}
