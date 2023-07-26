package main

import "fmt"

func checkPrecisely(grid [][]int, k, m int) bool {
	for i := 0; i < len(grid); i++ {
		if grid[k][i] != grid[i][m] {
			return false
		}
	}
	return true
}

func equalPairs(grid [][]int) int {
	row_map := make(map[int][]int)
	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
		row_map[sum] = append(row_map[sum], i)
	}

	column_map := make(map[int][]int)
	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[j][i]
		}
		column_map[sum] = append(column_map[sum], i)
	}

	result := 0
	for key, value := range row_map {
		if column_value, ok := column_map[key]; ok {
			for i := 0; i < len(value); i++ {
				for j := 0; j < len(column_value); j++ {
					if checkPrecisely(grid, value[i], column_value[j]) {
						result++
					}
				}
			}
		}
	}

	return result
}

func main() {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 4}, {2, 4, 2, 2}, {2, 5, 2, 2}}
	fmt.Println(equalPairs(grid))
}
