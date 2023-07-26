package main

import (
	"fmt"
	"strconv"
)

func createRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}

	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}

func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	n := len(nums)
	if n == 0 {
		return result
	}

	i := nums[0]
	current_start := nums[0]
	current_end := nums[0]
	for i = 1; i < n; i++ {
		if nums[i]-current_end > 1 {
			result = append(result, createRange(current_start, current_end))
			current_start = nums[i]
			current_end = nums[i]
		}
		current_end = nums[i]
	}

	result = append(result, createRange(current_start, current_end))

	return result
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
