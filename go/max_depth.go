package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, depth int) int {
	var left_depth, right_depth int = depth, 0
	if root.Left != nil {
		left_depth = traverse(root.Left, depth+1)
	}
	if root.Right != nil {
		right_depth = traverse(root.Right, depth+1)
	}
	if left_depth > right_depth {
		return left_depth
	} else {
		return right_depth
	}
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 1)
}

func main() {
	root := &TreeNode{}
	fmt.Println(maxDepth(root))
}
