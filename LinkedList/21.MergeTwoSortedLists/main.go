package main

import (
	"fmt"
)

// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:

// Input: list1 = [], list2 = []
// Output: []

// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	result := new(ListNode)
	dummy := result

	ptr1, ptr2 := list1, list2

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			dummy.Next = ptr1
			dummy = dummy.Next
			ptr1 = ptr1.Next
		} else {
			dummy.Next = ptr2
			dummy = dummy.Next
			ptr2 = ptr2.Next
		}
	}

	for ptr1 != nil {
		dummy.Next = ptr1
		dummy = dummy.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		dummy.Next = ptr2
		dummy = dummy.Next
		ptr2 = ptr2.Next
	}

	result = result.Next
	return result
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(mergeTwoLists(head, head1)) // [1,1,2,3,4,4]
}
