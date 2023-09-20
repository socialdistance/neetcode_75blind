package main

import "fmt"

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Input: root = [3,9,20,null,null,15,7]
// Output: 3

// Example 2:

// Input: root = [1,null,2]
// Output: 2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) insert(id *TreeNode) {
	current := t

	if id.Val <= current.Val {
		if current.Left == nil {
			current.Left = &TreeNode{Val: id.Val}
		} else {
			current.Left.insert(id)
		}
	} else {
		if current.Right == nil {
			current.Right = &TreeNode{Val: id.Val}
		} else {
			current.Right.insert(id)
		}
	}
}

func (t *TreeNode) inOrder(root *TreeNode) {
	if root != nil {
		fmt.Println("data:", root.Val)
		root.inOrder(root.Left)
		root.inOrder(root.Right)
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func main() {
	root := &TreeNode{
		Val: 3,
	}

	root.insert(&TreeNode{Val: 9})
	root.insert(&TreeNode{Val: 20})
	root.insert(&TreeNode{})
	root.insert(&TreeNode{})
	root.insert(&TreeNode{Val: 15})
	root.insert(&TreeNode{Val: 7})

	root.inOrder(root)

	fmt.Println(maxDepth(root))
	// 3,9,20,null,null,15,7
}
