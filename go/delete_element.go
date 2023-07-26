package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	nums = nums[0:index]
	return index
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeElement(arr, 2))
	fmt.Println(arr)
}
