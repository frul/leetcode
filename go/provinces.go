package main

import "fmt"

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(arr []int) int {
	if len(arr) == 0 {
		return 0 // or any other default value
	}
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	union := make([]int, size)
	for i := range union {
		union[i] = -1
	}
	latest_index := 0
	for i := 0; i < len(isConnected); i++ {
		var index int
		var old_indices []int
		if union[i] != -1 {
			index = union[i]
		} else {
			index = latest_index
			latest_index++
		}
		for j := 0; j < len(isConnected[i]); j++ {
			if i == j {
				continue
			}
			if isConnected[i][j] == 1 {
				if union[j] != -1 {
					old_indices = append(old_indices, union[j])
				}
			}
		}

		index = min(index, findMin(old_indices))
		for j := 0; j < len(union); j++ {
			if contains(old_indices, union[j]) {
				union[j] = index
			}
		}
	}

	unique := make(map[int]bool)
	for _, val := range union {
		unique[val] = true
	}
	return len(unique)
}

func main() {
	matrix := [][]int{
		{1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0}, {1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}}
	fmt.Println(findCircleNum(matrix))
}
