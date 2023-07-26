package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2
	for left < right {
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
		if mid < 1 {
			mid = 1
		} else if mid > len(arr)-1 {
			mid = len(arr) - 2
		}
	}
	return mid
}
func main() {
	nums := []int{3, 5, 3, 2, 0}
	fmt.Println(peakIndexInMountainArray(nums))
}
