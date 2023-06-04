package main

import "fmt"

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func square(a int, b int, c int, d int) int {
	return min(a, b) * (d - c)
}

func maxArea(height []int) int {
    var result int
	var left int = 0
	var right int = len(height) - 1
	for left < right {
		var sq = square(height[left], height[right], left, right)
		if sq > result {
			result = sq
		}
		if (height[left + 1] > height[left]) {
			left++
		} else if (height[right - 1] > height[right]) {
			right--
		} else {
			left++
			right--
		}
	}
	return result
}

func main() {
	arr := []int{1,3,2,5,25,24,5}
    fmt.Println(maxArea(arr))
}