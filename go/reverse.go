package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		the_next := curr.Next
		curr.Next = prev
		prev = curr
		curr = the_next
	}
	return prev
}

func main() {
	fmt.Println("Hi!")
}
