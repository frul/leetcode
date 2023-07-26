package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var prev, slow, fast *ListNode = nil, head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		} else {
			fast = fast.Next
		}
		prev = slow
		slow = slow.Next
	}

	prev.Next = slow.Next
	return head
}

func main() {

}
