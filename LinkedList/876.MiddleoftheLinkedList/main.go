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

var tempNode = &ListNode{
	Val:  -1,
	Next: nil,
}

var current = tempNode

var count = 0

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tempNode := new(ListNode)

	current := tempNode

	if head.Next != nil {
		count++
		middleNode(head.Next)
	}

	current.Next = head
	current = current.Next

	for count/2 > 0 {
		count--
	}

	fmt.Println(tempNode.Next)

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

	fmt.Println(middleNode(head)) // [3,4,5]
}
