package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	in1 := make(map[int]struct{})
	in2 := make(map[int]struct{})
	for _, value := range nums1 {
		in1[value] = struct{}{}
	}
	for _, value := range nums2 {
		in2[value] = struct{}{}
	}

	result := make([][]int, 2)
	for key, _ := range in1 {
		if _, ok := in2[key]; !ok {
			result[0] = append(result[0], key)
		}
	}
	for key, _ := range in2 {
		if _, ok := in1[key]; !ok {
			result[1] = append(result[1], key)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}
