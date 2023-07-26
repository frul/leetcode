package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func topKFrequent(nums []int, k int) []int {
	freq_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq_map[nums[i]] += 1
	}

	keys_of_map := make([]int, len(freq_map))
	for key, _ := range freq_map {
		keys_of_map = append(keys_of_map, key)
	}

	sort.Slice(keys_of_map, func(i, j int) bool {
		return freq_map[keys_of_map[i]] > freq_map[keys_of_map[j]]
	})

	size := min(k, len(keys_of_map))
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = keys_of_map[i]
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
