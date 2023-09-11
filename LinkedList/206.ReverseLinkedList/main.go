package main

import (
	"fmt"
)

// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) display(head *ListNode) {
	fmt.Println("List: first -> last:", head)
	current := head

	for current != nil {
		current.display(current.Next)
		current = current.Next
	}
}

var tempNode = &ListNode{
	Val:  -1,
	Next: nil,
}

var current = tempNode

func reverseList(head *ListNode) *ListNode {
	// var tempNode, current *ListNode

	if head == nil {
		return head
	}

	if head.Next != nil {
		reverseList(head.Next)
	}

	current.Next = head
	current = current.Next
	current.Next = nil

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
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(reverseList(head))
	// head.display(head)
}
