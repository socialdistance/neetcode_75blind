package main

import "fmt"

// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln

// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

// Input: head = [1,2,3,4]
// Output: [1,4,2,3]

// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]

type ListNode struct {
	Val  int
	Next *ListNode
}

var tempNode = new(ListNode)
var tmp = tempNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	reversed := reverse(slow.Next)

	fmt.Println(reversed)

}

func reverse(node *ListNode) *ListNode {
	if node == nil {
		return node
	}

	tempNode = new(ListNode)
	tmp = tempNode

	if node.Next != nil {
		reverse(node.Next)
	}

	tmp.Next = node
	tmp = tmp.Next
	tmp.Next = nil

	return tempNode.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	reorderList(head) // 1,4,2,3
}
