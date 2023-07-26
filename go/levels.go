package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverse(root *TreeNode, level int, sums map[int]int) {
	if root == nil {
		return
	}
	traverse(root.Right, level+1, sums)
	sums[level] += root.Val
	traverse(root.Left, level+1, sums)
}

func maxLevelSum(root *TreeNode) int {
	sums := make(map[int]int)
	traverse(root, 1, sums)
	max_val := math.MinInt32
	for _, value := range sums {
		max_val = max(max_val, value)
	}
	result := len(sums) + 1
	for key, value := range sums {
		if value == max_val {
			if key < result {
				result = key
			}
		}
	}
	return result
}

func buildTree(nums []interface{}, index int) *TreeNode {
	if index >= len(nums) || nums[index] == nil {
		return nil
	}
	root := &TreeNode{Val: nums[index].(int)}
	root.Left = buildTree(nums, 2*index+1)
	root.Right = buildTree(nums, 2*index+2)
	return root
}

func main() {
	nums := []interface{}{1, 7, 0, 7, -8, nil, nil}
	root := buildTree(nums, 0)
	fmt.Println(maxLevelSum(root))
}
