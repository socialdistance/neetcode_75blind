package main

import (
	"fmt"
)

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) display() {
	fmt.Println(l)
	h := l
	for h.Next != nil {
		h.display()
		h = h.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head

	for current.Next != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}

		if current.Next != nil {
			current = current.Next
		}
	}

	// current.display()

	return head
}

func main() {
	_ = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(head)) // 1, 2
}
