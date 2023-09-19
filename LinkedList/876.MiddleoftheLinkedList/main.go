package main

import "fmt"

// Given the head of a singly linked list, return the middle node of the linked list.

// If there are two middle nodes, return the second middle node.

// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.

// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

type ListNode struct {
	Val  int
	Next *ListNode
}

var tmpNode = new(ListNode)

var count = 0
var middle = 0

func middleNode(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
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
					// Next: &ListNode{
					// 	Val:  5,
					// 	Next: nil,
					// },
				},
			},
		},
	}

	fmt.Println(middleNode(head)) // [3,4,5]
}
