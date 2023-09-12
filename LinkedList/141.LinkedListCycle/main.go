package main

import "fmt"

// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

type ListListNode struct {
	Val  int
	Next *ListListNode
}

func hasCycle(head *ListListNode) bool {
	// cycle := make(map[*ListListNode]bool)

	// for head != nil {
	// 	if head == nil || head.Next == nil {
	// 		return false
	// 	}

	// 	if _, ok := cycle[head.Next]; ok {
	// 		return true
	// 	}

	// 	cycle[head] = false

	// 	head = head.Next
	// }

	// return false

	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	head := &ListListNode{Val: 1} //initialize the linked list with Vals
	ListNode2 := &ListListNode{Val: 2}
	ListNode3 := &ListListNode{Val: 3}
	ListNode4 := &ListListNode{Val: 4}

	head.Next = ListNode2 //point to the elements using linked list
	ListNode2.Next = ListNode3
	ListNode3.Next = ListNode4
	ListNode4.Next = ListNode2

	fmt.Println(hasCycle(head)) // Output: true
}
