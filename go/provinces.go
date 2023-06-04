package main

import "fmt"

func traverse(isConnected [][]int, visited []bool, i int) {
	visited[i] = true
	var stack []int
	for j := 0; j < len(isConnected); j++ {
		if i == j {
			continue
		}
		if isConnected[i][j] == 1 {
			if !visited[j] {
				stack = append(stack, j)
			}
		}
	}

	for len(stack) > 0 {
		n := len(stack) - 1
		j := stack[n]
		traverse(isConnected, visited, j)
		stack = stack[:n] // pop
	}
}

func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	visited := make([]bool, size)
	for i := range visited {
		visited[i] = false
	}

	var count = 0
	for i := 0; i < len(isConnected); i++ {
		if visited[i] == true {
			continue
		}
		count++
		traverse(isConnected, visited, i)
	}
	return count
}

func main() {
	matrix := [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(matrix))
}
