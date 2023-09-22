package main

import "fmt"

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) inOrder(root *TreeNode) {
	if root != nil {
		fmt.Println("data:", root.Val)
		root.inOrder(root.Left)
		root.inOrder(root.Right)
	}

}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// p = [1,2,3], q = [1,2,3]
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val:   2,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   3,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// }

	// root1 := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val:   2,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// 	Right: &TreeNode{
	// 		Val:   3,
	// 		Left:  nil,
	// 		Right: nil,
	// 	},
	// }

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	root3 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isSameTree(root2, root3))

}
